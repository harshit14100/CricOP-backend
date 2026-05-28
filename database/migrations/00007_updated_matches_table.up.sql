ALTER TABLE matches
    ADD COLUMN batting_team_id UUID REFERENCES teams(team_id),
    ADD COLUMN bowling_team_id UUID REFERENCES teams(team_id),
    ADD COLUMN striker_id UUID REFERENCES users(id),
    ADD COLUMN non_striker_id UUID REFERENCES users(id),
    ADD COLUMN current_bowler_id UUID REFERENCES users(id),
    ADD COLUMN current_inning INT DEFAULT 1,
    ADD COLUMN total_runs INT DEFAULT 0,
    ADD COLUMN total_wickets INT DEFAULT 0,
    ADD COLUMN current_over INT DEFAULT 0,
    ADD COLUMN current_ball INT DEFAULT 0;