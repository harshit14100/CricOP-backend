ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_inning_id_fkey
        FOREIGN KEY (inning_id)
            REFERENCES innings(id);

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_striker_id_fkey
        FOREIGN KEY (striker_id)
            REFERENCES players(id);

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_non_striker_id_fkey
        FOREIGN KEY (non_striker_id)
            REFERENCES players(id);

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_bowler_id_fkey
        FOREIGN KEY (bowler_id)
            REFERENCES players(id);

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_fielder_id_fkey
        FOREIGN KEY (fielder_id)
            REFERENCES players(id);

ALTER TABLE deliveries
    ADD CONSTRAINT deliveries_player_out_id_fkey
        FOREIGN KEY (player_out_id)
            REFERENCES players(id);