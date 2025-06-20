--
-- PostgreSQL database dump
--

-- Dumped from database version 15.6 (Debian 15.6-1.pgdg120+2)
-- Dumped by pg_dump version 15.6 (Debian 15.6-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA public;


--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- Name: check_unique_player_number(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.check_unique_player_number() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
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
	$$;


--
-- Name: log_user_login(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.log_user_login() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
	BEGIN
		INSERT INTO user_login_logs (user_id, token, time)
		VALUES (NEW.user_id, NEW.token, NOW());
		RETURN NEW;
	END;
	$$;


--
-- Name: set_token_expiry(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.set_token_expiry() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
	BEGIN
		IF NEW.expires_at IS NULL THEN
			NEW.expires_at := NOW() + INTERVAL '7 days';
		END IF;
		RETURN NEW;
	END;
	$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: action_rates; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.action_rates (
    action_rate_id bigint NOT NULL,
    help_text text,
    signature text,
    action_id bigint
);


--
-- Name: action_rates_action_rate_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.action_rates_action_rate_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: action_rates_action_rate_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.action_rates_action_rate_id_seq OWNED BY public.action_rates.action_rate_id;


--
-- Name: actions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.actions (
    action_id bigint NOT NULL,
    name text
);


--
-- Name: actions_action_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.actions_action_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: actions_action_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.actions_action_id_seq OWNED BY public.actions.action_id;


--
-- Name: ampluas; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.ampluas (
    amplua_id bigint NOT NULL,
    name text
);


--
-- Name: ampluas_amplua_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.ampluas_amplua_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: ampluas_amplua_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.ampluas_amplua_id_seq OWNED BY public.ampluas.amplua_id;


--
-- Name: auth_tokens; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.auth_tokens (
    token text NOT NULL,
    user_id bigint NOT NULL,
    expires_at timestamp with time zone
);


--
-- Name: championships; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.championships (
    championship_id bigint NOT NULL,
    title text,
    start_date timestamp with time zone,
    end_date timestamp with time zone
);


--
-- Name: championships_championship_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.championships_championship_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: championships_championship_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.championships_championship_id_seq OWNED BY public.championships.championship_id;


--
-- Name: games; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.games (
    game_id bigint NOT NULL,
    date timestamp with time zone,
    win boolean,
    round_id bigint,
    team_id bigint,
    opp_team_id bigint,
    team_score bigint,
    opp_team_score bigint
);


--
-- Name: games_game_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.games_game_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: games_game_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.games_game_id_seq OWNED BY public.games.game_id;


--
-- Name: players; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.players (
    player_id bigint NOT NULL,
    first_name text,
    last_name text,
    birthdate timestamp with time zone,
    gender text,
    height numeric,
    number bigint,
    team_id bigint,
    amplua_id bigint,
    CONSTRAINT chk_players_gender CHECK ((gender = ANY (ARRAY['male'::text, 'female'::text])))
);


--
-- Name: players_player_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.players_player_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: players_player_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.players_player_id_seq OWNED BY public.players.player_id;


--
-- Name: rounds; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.rounds (
    round_id bigint NOT NULL,
    serial_number smallint,
    start_date timestamp with time zone,
    end_date timestamp with time zone,
    championship_id bigint
);


--
-- Name: rounds_round_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.rounds_round_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: rounds_round_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.rounds_round_id_seq OWNED BY public.rounds.round_id;


--
-- Name: set_actions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.set_actions (
    set_action_id bigint NOT NULL,
    set_id bigint,
    player_id bigint,
    action_rate_id bigint
);


--
-- Name: set_actions_set_action_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.set_actions_set_action_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: set_actions_set_action_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.set_actions_set_action_id_seq OWNED BY public.set_actions.set_action_id;


--
-- Name: sets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sets (
    set_id bigint NOT NULL,
    serial_number bigint,
    team_score bigint,
    opp_score bigint,
    game_id bigint
);


--
-- Name: sets_set_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.sets_set_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sets_set_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.sets_set_id_seq OWNED BY public.sets.set_id;


--
-- Name: teams; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.teams (
    team_id bigint NOT NULL,
    name text,
    user_id bigint
);


--
-- Name: teams_team_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.teams_team_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: teams_team_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.teams_team_id_seq OWNED BY public.teams.team_id;


--
-- Name: user_login_logs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_login_logs (
    id bigint NOT NULL,
    user_id bigint,
    token text,
    "time" timestamp with time zone
);


--
-- Name: user_login_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_login_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: user_login_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_login_logs_id_seq OWNED BY public.user_login_logs.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    user_id bigint NOT NULL,
    name text,
    password text,
    email text NOT NULL
);


--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: action_rates action_rate_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.action_rates ALTER COLUMN action_rate_id SET DEFAULT nextval('public.action_rates_action_rate_id_seq'::regclass);


--
-- Name: actions action_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.actions ALTER COLUMN action_id SET DEFAULT nextval('public.actions_action_id_seq'::regclass);


--
-- Name: ampluas amplua_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ampluas ALTER COLUMN amplua_id SET DEFAULT nextval('public.ampluas_amplua_id_seq'::regclass);


--
-- Name: championships championship_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.championships ALTER COLUMN championship_id SET DEFAULT nextval('public.championships_championship_id_seq'::regclass);


--
-- Name: games game_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games ALTER COLUMN game_id SET DEFAULT nextval('public.games_game_id_seq'::regclass);


--
-- Name: players player_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players ALTER COLUMN player_id SET DEFAULT nextval('public.players_player_id_seq'::regclass);


--
-- Name: rounds round_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rounds ALTER COLUMN round_id SET DEFAULT nextval('public.rounds_round_id_seq'::regclass);


--
-- Name: set_actions set_action_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.set_actions ALTER COLUMN set_action_id SET DEFAULT nextval('public.set_actions_set_action_id_seq'::regclass);


--
-- Name: sets set_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sets ALTER COLUMN set_id SET DEFAULT nextval('public.sets_set_id_seq'::regclass);


--
-- Name: teams team_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.teams ALTER COLUMN team_id SET DEFAULT nextval('public.teams_team_id_seq'::regclass);


--
-- Name: user_login_logs id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_login_logs ALTER COLUMN id SET DEFAULT nextval('public.user_login_logs_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: action_rates; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.action_rates (action_rate_id, help_text, signature, action_id) FROM stdin;
1	Очко на атаке	++	3
2	Обратно переводят без атаки	+-	3
3	Обратно переводят с атакой	_	3
4	Ошибка на подаче	-	3
5	Поднял в защите	++	4
6	Ошибка в защите	-	4
7	Идеальный прием между 2 и 3 зоной	++	5
8	Прием в пределах 3-хметровой линии	+	5
9	Прием вверх	_	5
10	Ошибка	-	5
11	Эйс	++	1
12	Обычная подача	_	1
13	Ошибка на подаче	-	1
14	Заработали очко на блоке	++	2
15	Ничего/Ошибка на блоке	-	2
\.


--
-- Data for Name: actions; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.actions (action_id, name) FROM stdin;
1	Подача
2	Блок
3	Атака
4	Защита
5	Прием
\.


--
-- Data for Name: ampluas; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.ampluas (amplua_id, name) FROM stdin;
1	Доигровщик
2	Центральный блокирующий
3	Либеро
4	Связующий
5	Диагональный
\.


--
-- Data for Name: auth_tokens; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.auth_tokens (token, user_id, expires_at) FROM stdin;
SjrQgqqNSjW4Ty09gswNf0EPqjwFHtrO	1	2025-06-27 07:49:42.081863+00
\.


--
-- Data for Name: championships; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.championships (championship_id, title, start_date, end_date) FROM stdin;
1	Summer Championship 2025	2025-07-01 00:00:00+00	2025-07-15 00:00:00+00
\.


--
-- Data for Name: games; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.games (game_id, date, win, round_id, team_id, opp_team_id, team_score, opp_team_score) FROM stdin;
1	2025-07-01 15:00:00+00	\N	1	1	2	\N	\N
2	2025-07-01 00:00:00+00	\N	1	2	6	\N	\N
3	2025-07-02 00:00:00+00	\N	1	5	1	\N	\N
4	2025-07-03 00:00:00+00	\N	1	2	5	\N	\N
5	2025-07-04 00:00:00+00	\N	1	2	1	\N	\N
\.


--
-- Data for Name: players; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.players (player_id, first_name, last_name, birthdate, gender, height, number, team_id, amplua_id) FROM stdin;
1	Александр	Поляков	2006-01-07 00:00:00+00	male	1.91	17	1	5
2	Имя2	Фамилия2	2006-01-01 00:00:00+00	male	1.92	2	1	2
3	Имя3	Фамилия3	2006-01-01 00:00:00+00	male	1.76	3	1	4
4	Имя4	Фамилия4	2006-01-01 00:00:00+00	male	1.91	4	1	2
5	Имя5	Фамилия5	2006-01-01 00:00:00+00	male	1.81	5	1	5
6	Имя6	Фамилия6	2006-01-01 00:00:00+00	male	1.95	6	1	4
7	Имя7	Фамилия7	2006-01-01 00:00:00+00	male	1.82	7	1	1
8	Имя8	Фамилия8	2006-01-01 00:00:00+00	male	1.93	8	1	4
9	Имя9	Фамилия9	2006-01-01 00:00:00+00	male	1.89	9	2	4
10	Имя10	Фамилия10	2006-01-01 00:00:00+00	male	2.01	10	2	1
11	Имя11	Фамилия11	2006-01-01 00:00:00+00	male	1.88	11	2	5
12	Имя12	Фамилия12	2006-01-01 00:00:00+00	male	2.04	12	2	5
13	Имя13	Фамилия13	2006-01-01 00:00:00+00	male	2.03	13	2	2
14	Имя14	Фамилия14	2006-01-01 00:00:00+00	male	2	14	2	1
15	Имя15	Фамилия15	2006-01-01 00:00:00+00	male	1.93	15	2	3
16	Имя16	Фамилия16	2006-01-01 00:00:00+00	male	1.95	16	5	5
17	Имя17	Фамилия17	2006-01-01 00:00:00+00	male	1.87	17	5	4
18	Имя18	Фамилия18	2006-01-01 00:00:00+00	male	2.03	18	5	5
19	Имя19	Фамилия19	2006-01-01 00:00:00+00	male	1.78	19	5	2
20	Имя20	Фамилия20	2006-01-01 00:00:00+00	male	1.79	20	5	3
21	Имя21	Фамилия21	2006-01-01 00:00:00+00	male	1.76	21	5	5
22	Имя22	Фамилия22	2006-01-01 00:00:00+00	male	2.02	22	5	3
23	Имя23	Фамилия23	2006-01-01 00:00:00+00	male	2.01	23	5	5
24	Имя24	Фамилия24	2006-01-01 00:00:00+00	male	1.8	24	6	2
25	Имя25	Фамилия25	2006-01-01 00:00:00+00	male	1.82	25	6	3
26	Имя26	Фамилия26	2006-01-01 00:00:00+00	male	1.89	26	6	1
27	Имя27	Фамилия27	2006-01-01 00:00:00+00	male	2	27	6	5
28	Имя28	Фамилия28	2006-01-01 00:00:00+00	male	1.77	28	6	5
29	Имя29	Фамилия29	2006-01-01 00:00:00+00	male	2.01	29	6	1
\.


--
-- Data for Name: rounds; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.rounds (round_id, serial_number, start_date, end_date, championship_id) FROM stdin;
1	1	2025-06-30 10:00:00+00	\N	1
\.


--
-- Data for Name: set_actions; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.set_actions (set_action_id, set_id, player_id, action_rate_id) FROM stdin;
1	1	1	1
2	1	1	1
3	1	1	3
4	1	1	4
5	1	1	3
6	1	1	3
7	1	1	3
8	1	1	1
9	1	1	1
10	1	1	1
11	1	1	4
12	1	1	4
13	3	11	10
14	3	11	13
15	3	12	11
16	3	12	14
17	3	12	7
18	3	12	5
19	3	10	13
20	3	10	4
21	3	10	1
22	3	10	15
23	3	13	6
24	3	13	13
25	3	14	15
26	3	14	15
27	3	14	8
28	4	12	11
29	4	12	14
30	4	12	4
31	4	15	15
32	4	15	5
33	4	15	9
34	4	15	5
35	4	10	14
36	4	10	8
37	4	10	12
38	4	10	9
39	4	11	5
40	4	11	10
41	4	11	10
42	4	11	9
43	4	9	5
44	4	9	15
45	4	9	7
46	4	9	9
47	5	10	13
48	5	10	8
49	5	10	9
50	5	10	6
51	5	15	12
52	5	15	8
53	5	9	5
54	5	9	15
55	5	9	14
56	5	14	14
57	5	14	5
58	5	14	8
59	5	14	14
60	5	13	10
61	5	13	8
62	6	12	5
63	6	12	5
64	6	12	8
65	6	11	7
66	6	11	15
67	6	11	3
68	6	11	11
69	6	14	3
70	6	14	5
71	6	13	5
72	6	13	10
73	6	13	6
74	6	9	12
75	6	9	5
76	6	9	14
77	7	21	11
78	7	21	8
79	7	19	5
80	7	19	12
81	7	19	14
82	7	19	12
83	7	23	7
84	7	23	15
85	7	23	2
86	7	23	6
87	7	16	5
88	7	16	12
89	7	16	4
90	7	20	5
91	7	20	10
92	7	20	5
93	8	23	7
94	8	23	15
95	8	23	14
96	8	23	9
97	8	20	12
98	8	20	15
99	8	20	6
100	8	20	11
101	8	1	12
102	8	1	3
103	8	1	12
104	8	1	4
105	8	22	5
106	8	22	11
107	8	22	15
108	8	22	10
109	8	21	3
110	8	21	9
111	9	19	5
112	9	19	9
113	9	22	14
114	9	22	1
115	9	18	9
116	9	18	15
117	9	18	5
118	9	18	15
119	9	1	9
120	9	1	9
121	9	23	14
122	9	23	13
123	10	20	2
124	10	20	9
125	10	20	15
126	10	22	2
127	10	22	15
128	10	16	9
129	10	16	2
130	10	16	15
131	10	16	13
132	10	23	13
133	10	23	14
134	10	23	8
135	10	23	7
136	10	1	10
137	10	1	15
138	10	1	5
139	10	1	5
140	11	22	2
141	11	22	2
142	11	19	7
143	11	19	15
144	11	1	14
145	11	1	13
146	11	1	6
147	11	21	5
148	11	21	4
149	11	21	6
150	11	21	6
151	11	18	2
152	11	18	11
153	12	9	6
154	12	9	2
155	12	9	6
156	12	9	11
157	12	11	13
158	12	11	13
159	12	11	13
160	12	12	10
161	12	12	4
162	12	12	13
163	12	12	14
164	12	15	15
165	12	15	1
166	12	15	4
167	12	15	12
168	12	10	7
169	12	10	12
170	12	10	10
171	13	11	8
172	13	11	8
173	13	11	6
174	13	11	12
175	13	14	6
176	13	14	13
177	13	14	1
178	13	12	3
179	13	12	15
180	13	12	1
181	13	15	3
182	13	15	9
183	13	15	14
184	13	15	11
185	13	9	5
186	13	9	3
187	14	15	2
188	14	15	2
189	14	15	3
190	14	15	10
191	14	13	3
192	14	13	11
193	14	13	11
194	14	13	8
195	14	9	14
196	14	9	12
197	14	9	11
198	14	9	13
199	14	10	8
200	14	10	15
201	14	10	5
202	14	11	3
203	14	11	5
204	14	11	12
205	14	11	2
206	15	15	7
207	15	15	10
208	15	11	14
209	15	11	6
210	15	11	5
211	15	13	11
212	15	13	9
213	15	13	2
214	15	14	9
215	15	14	14
216	15	14	8
217	15	14	2
218	15	12	15
219	15	12	15
220	15	12	6
221	16	10	14
222	16	10	6
223	16	15	12
224	16	15	2
225	16	15	5
226	16	14	15
227	16	14	11
228	16	14	2
229	16	14	11
230	16	12	13
231	16	12	12
232	16	12	5
233	16	13	7
234	16	13	6
235	16	13	2
236	17	14	13
237	17	14	12
238	17	11	1
239	17	11	7
240	17	10	4
241	17	10	15
242	17	9	4
243	17	9	13
244	17	13	4
245	17	13	6
246	17	13	5
247	17	13	14
248	18	11	8
249	18	11	14
250	18	14	9
251	18	14	3
252	18	14	6
253	18	12	15
254	18	12	5
255	18	15	9
256	18	15	7
257	18	15	12
258	18	15	5
259	18	9	4
260	18	9	9
261	18	9	6
262	19	12	15
263	19	12	1
264	19	12	5
265	19	12	3
266	19	11	6
267	19	11	13
268	19	11	4
269	19	9	6
270	19	9	9
271	19	9	14
272	19	9	14
273	19	10	6
274	19	10	1
275	19	10	13
276	19	15	13
277	19	15	11
278	19	15	1
279	19	15	1
280	20	12	13
281	20	12	8
282	20	12	3
283	20	9	11
284	20	9	6
285	20	9	15
286	20	11	14
287	20	11	15
288	20	13	12
289	20	13	2
290	20	10	15
291	20	10	6
\.


--
-- Data for Name: sets; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.sets (set_id, serial_number, team_score, opp_score, game_id) FROM stdin;
1	1	\N	\N	1
2	2	\N	\N	1
3	1	\N	\N	2
4	2	\N	\N	2
5	3	\N	\N	2
6	4	\N	\N	2
7	1	\N	\N	3
8	2	\N	\N	3
9	3	\N	\N	3
10	4	\N	\N	3
11	5	\N	\N	3
12	1	\N	\N	4
13	2	\N	\N	4
14	3	\N	\N	4
15	4	\N	\N	4
16	5	\N	\N	4
17	1	\N	\N	5
18	2	\N	\N	5
19	3	\N	\N	5
20	4	\N	\N	5
\.


--
-- Data for Name: teams; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.teams (team_id, name, user_id) FROM stdin;
1	Кронверские Барсы	1
2	ГУТ	1
3	Кронверские Барсы	1
4	Кронверские Барсы	1
5	Кронверские Барсы	1
6	Северные Волки	1
\.


--
-- Data for Name: user_login_logs; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.user_login_logs (id, user_id, token, "time") FROM stdin;
1	1	SjrQgqqNSjW4Ty09gswNf0EPqjwFHtrO	2025-06-20 07:49:42.082599+00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (user_id, name, password, email) FROM stdin;
1	user	$2a$10$mvFO0nr1PKsWMJCQO.m6Qe87cEcVDqRnxUfMS0KKiRIUeMxUPvW2q	user@mail.ru
\.


--
-- Name: action_rates_action_rate_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.action_rates_action_rate_id_seq', 15, true);


--
-- Name: actions_action_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.actions_action_id_seq', 5, true);


--
-- Name: ampluas_amplua_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.ampluas_amplua_id_seq', 5, true);


--
-- Name: championships_championship_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.championships_championship_id_seq', 1, true);


--
-- Name: games_game_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.games_game_id_seq', 5, true);


--
-- Name: players_player_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.players_player_id_seq', 29, true);


--
-- Name: rounds_round_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.rounds_round_id_seq', 1, true);


--
-- Name: set_actions_set_action_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.set_actions_set_action_id_seq', 291, true);


--
-- Name: sets_set_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.sets_set_id_seq', 20, true);


--
-- Name: teams_team_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.teams_team_id_seq', 6, true);


--
-- Name: user_login_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.user_login_logs_id_seq', 1, true);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_user_id_seq', 1, true);


--
-- Name: action_rates action_rates_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.action_rates
    ADD CONSTRAINT action_rates_pkey PRIMARY KEY (action_rate_id);


--
-- Name: actions actions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT actions_pkey PRIMARY KEY (action_id);


--
-- Name: ampluas ampluas_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.ampluas
    ADD CONSTRAINT ampluas_pkey PRIMARY KEY (amplua_id);


--
-- Name: championships championships_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.championships
    ADD CONSTRAINT championships_pkey PRIMARY KEY (championship_id);


--
-- Name: games games_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games
    ADD CONSTRAINT games_pkey PRIMARY KEY (game_id);


--
-- Name: players players_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_pkey PRIMARY KEY (player_id);


--
-- Name: rounds rounds_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rounds
    ADD CONSTRAINT rounds_pkey PRIMARY KEY (round_id);


--
-- Name: set_actions set_actions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.set_actions
    ADD CONSTRAINT set_actions_pkey PRIMARY KEY (set_action_id);


--
-- Name: sets sets_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sets
    ADD CONSTRAINT sets_pkey PRIMARY KEY (set_id);


--
-- Name: teams teams_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.teams
    ADD CONSTRAINT teams_pkey PRIMARY KEY (team_id);


--
-- Name: auth_tokens uni_auth_tokens_token; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auth_tokens
    ADD CONSTRAINT uni_auth_tokens_token UNIQUE (token);


--
-- Name: users uni_users_email; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_email UNIQUE (email);


--
-- Name: user_login_logs user_login_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_login_logs
    ADD CONSTRAINT user_login_logs_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: idx_set_actions_action_rate_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_set_actions_action_rate_id ON public.set_actions USING btree (action_rate_id);


--
-- Name: idx_set_actions_player_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_set_actions_player_id ON public.set_actions USING btree (player_id);


--
-- Name: idx_set_actions_set_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_set_actions_set_id ON public.set_actions USING btree (set_id);


--
-- Name: auth_tokens set_token_expiry_trigger; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER set_token_expiry_trigger BEFORE INSERT ON public.auth_tokens FOR EACH ROW EXECUTE FUNCTION public.set_token_expiry();


--
-- Name: auth_tokens trg_log_user_login; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trg_log_user_login AFTER INSERT ON public.auth_tokens FOR EACH ROW EXECUTE FUNCTION public.log_user_login();


--
-- Name: players trigger_check_unique_player_number; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER trigger_check_unique_player_number BEFORE INSERT OR UPDATE ON public.players FOR EACH ROW EXECUTE FUNCTION public.check_unique_player_number();


--
-- Name: action_rates fk_actions_rates; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.action_rates
    ADD CONSTRAINT fk_actions_rates FOREIGN KEY (action_id) REFERENCES public.actions(action_id);


--
-- Name: players fk_ampluas_players; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT fk_ampluas_players FOREIGN KEY (amplua_id) REFERENCES public.ampluas(amplua_id);


--
-- Name: auth_tokens fk_auth_tokens_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.auth_tokens
    ADD CONSTRAINT fk_auth_tokens_user FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: rounds fk_championships_rounds; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rounds
    ADD CONSTRAINT fk_championships_rounds FOREIGN KEY (championship_id) REFERENCES public.championships(championship_id);


--
-- Name: games fk_games_opp_team; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games
    ADD CONSTRAINT fk_games_opp_team FOREIGN KEY (opp_team_id) REFERENCES public.teams(team_id);


--
-- Name: sets fk_games_sets; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sets
    ADD CONSTRAINT fk_games_sets FOREIGN KEY (game_id) REFERENCES public.games(game_id);


--
-- Name: games fk_games_team; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games
    ADD CONSTRAINT fk_games_team FOREIGN KEY (team_id) REFERENCES public.teams(team_id);


--
-- Name: games fk_rounds_games; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games
    ADD CONSTRAINT fk_rounds_games FOREIGN KEY (round_id) REFERENCES public.rounds(round_id);


--
-- Name: set_actions fk_set_actions_action_rate; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.set_actions
    ADD CONSTRAINT fk_set_actions_action_rate FOREIGN KEY (action_rate_id) REFERENCES public.action_rates(action_rate_id);


--
-- Name: set_actions fk_set_actions_player; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.set_actions
    ADD CONSTRAINT fk_set_actions_player FOREIGN KEY (player_id) REFERENCES public.players(player_id);


--
-- Name: set_actions fk_sets_set_actions; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.set_actions
    ADD CONSTRAINT fk_sets_set_actions FOREIGN KEY (set_id) REFERENCES public.sets(set_id);


--
-- Name: players fk_teams_players; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT fk_teams_players FOREIGN KEY (team_id) REFERENCES public.teams(team_id);


--
-- Name: user_login_logs fk_user_login_logs_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_login_logs
    ADD CONSTRAINT fk_user_login_logs_user FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: teams fk_users_teams; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.teams
    ADD CONSTRAINT fk_users_teams FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- PostgreSQL database dump complete
--

