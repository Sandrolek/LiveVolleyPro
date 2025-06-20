package db

import (
	"log"
	"volley/internal/db/seeds"
	"volley/internal/models/orm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=secret dbname=volley port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	const DROP = true

	if DROP {
		err = DB.Migrator().DropTable(
			&orm.SetAction{},
			&orm.ActionRate{},
			&orm.Action{},
			&orm.Set{},
			&orm.Game{},
			&orm.Round{},
			&orm.Championship{},
			&orm.Player{},
			&orm.Team{},
			&orm.Amplua{},
			&orm.User{},
			&orm.AuthToken{},
			&orm.UserLoginLog{},
		)
		if err != nil {
			log.Fatalf("Failed to drop tables: %v", err)
		}
	}

	err = DB.AutoMigrate(
		&orm.User{},
		&orm.Amplua{},
		&orm.Team{},
		&orm.Player{},
		&orm.Game{},
		&orm.Round{},
		&orm.Championship{},
		&orm.Set{},
		&orm.SetAction{},
		&orm.ActionRate{},
		&orm.Action{},
		&orm.AuthToken{},
		&orm.UserLoginLog{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	seeds.SeedAll(DB)

	createTokenExpiryTrigger := `
	CREATE OR REPLACE FUNCTION set_token_expiry()
	RETURNS TRIGGER AS $$
	BEGIN
		IF NEW.expires_at IS NULL THEN
			NEW.expires_at := NOW() + INTERVAL '7 days';
		END IF;
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	DROP TRIGGER IF EXISTS set_token_expiry_trigger ON auth_tokens;

	CREATE TRIGGER set_token_expiry_trigger
	BEFORE INSERT ON auth_tokens
	FOR EACH ROW
	EXECUTE FUNCTION set_token_expiry();
	`

	if err := DB.Exec(createTokenExpiryTrigger).Error; err != nil {
		log.Fatalf("Failed to create trigger and function: %v", err)
	}

	checkUniquePlayerNumber := `
	CREATE OR REPLACE FUNCTION check_unique_player_number()
	RETURNS TRIGGER AS $$
	BEGIN
	  IF EXISTS (
		SELECT 1 FROM players
		WHERE team_id = NEW.team_id
		  AND number = NEW.number
		  AND player_id != NEW.player_id
	  ) THEN
		RAISE EXCEPTION 'Player number % is already taken in team %', NEW.number, NEW.team_id;
	  END IF;
	  RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;
	
	DROP TRIGGER IF EXISTS trigger_check_unique_player_number ON players;
	
	CREATE TRIGGER trigger_check_unique_player_number
	BEFORE INSERT OR UPDATE ON players
	FOR EACH ROW
	EXECUTE FUNCTION check_unique_player_number();
	`

	if err := DB.Exec(checkUniquePlayerNumber).Error; err != nil {
		log.Fatalf("Failed to create player number uniqueness trigger: %v", err)
	}

	logUserLoginTrigger := `
	CREATE OR REPLACE FUNCTION log_user_login()
	RETURNS TRIGGER AS $$
	BEGIN
		INSERT INTO user_login_logs (user_id, token, time)
		VALUES (NEW.user_id, NEW.token, NOW());
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

	DROP TRIGGER IF EXISTS trg_log_user_login ON auth_tokens;

	CREATE TRIGGER trg_log_user_login
	AFTER INSERT ON auth_tokens
	FOR EACH ROW
	EXECUTE FUNCTION log_user_login();
	`

	if err := DB.Exec(logUserLoginTrigger).Error; err != nil {
		log.Fatalf("Failed to create user login logging trigger: %v", err)
	}

}
