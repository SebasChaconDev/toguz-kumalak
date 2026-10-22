CREATE TABLE synthetic_games(id text PRIMARY KEY, owner text NOT NULL, revision bigint NOT NULL CHECK(revision>=0), effects bigint NOT NULL CHECK(effects>=0));
CREATE TABLE synthetic_receipts(game_id text REFERENCES synthetic_games(id), actor text NOT NULL, operation_id text NOT NULL, intent_hash text NOT NULL, revision bigint NOT NULL, effects bigint NOT NULL, PRIMARY KEY(game_id,actor,operation_id));
INSERT INTO synthetic_games VALUES('owned-game','alice',0,0),('foreign-game','bob',0,0);
