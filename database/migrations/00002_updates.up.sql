ALTER TABLE player_stats
    RENAME COLUMN user_id TO player_id;

ALTER TABLE player_stats
    DROP CONSTRAINT player_stats_user_id_key;

ALTER TABLE player_stats
    DROP CONSTRAINT player_stats_user_id_fkey;

ALTER TABLE player_stats
    ADD CONSTRAINT player_stats_player_id_fkey
        FOREIGN KEY (player_id)
            REFERENCES players(id);

ALTER TABLE team_players
    DROP CONSTRAINT team_players_player_id_fkey;

ALTER TABLE team_players
    ADD CONSTRAINT team_players_player_id_fkey
        FOREIGN KEY (player_id)
            REFERENCES players(id);

ALTER TABLE player_stats
    DROP CONSTRAINT player_stats_player_id_fkey;

ALTER TABLE team_players
    ADD CONSTRAINT team_players_player_id_fkey
        FOREIGN KEY (player_id)
            REFERENCES players(id);

ALTER TABLE player_stats
    ADD CONSTRAINT player_stats_player_id_fkey
        FOREIGN KEY (player_id)
            REFERENCES players(id);

ALTER TABLE team_players
    DROP CONSTRAINT team_players_player_id_fkey;

ALTER TABLE team_players
    ADD CONSTRAINT team_players_player_id_fkey
        FOREIGN KEY (player_id)
            REFERENCES players(id);

ALTER TABLE matches
    DROP CONSTRAINT matches_man_of_match_id_fkey;

ALTER TABLE matches
    DROP CONSTRAINT matches_worst_player_id_fkey;

ALTER TABLE matches
    ADD CONSTRAINT matches_man_of_match_id_fkey
        FOREIGN KEY (man_of_match_id)
            REFERENCES players(id);

ALTER TABLE matches
    ADD CONSTRAINT matches_worst_player_id_fkey
        FOREIGN KEY (worst_player_id)
            REFERENCES players(id);

ALTER TABLE matches
    DROP CONSTRAINT matches_man_of_match_id_fkey;

ALTER TABLE matches
    DROP CONSTRAINT matches_worst_player_id_fkey;

ALTER TABLE matches
    ADD CONSTRAINT matches_man_of_match_id_fkey
        FOREIGN KEY (man_of_match_id)
            REFERENCES players(id);

ALTER TABLE matches
    ADD CONSTRAINT matches_worst_player_id_fkey
        FOREIGN KEY (worst_player_id)
            REFERENCES players(id);

ALTER TABLE match_players
    DROP CONSTRAINT match_players_player_id_fkey;