CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE match_status AS ENUM (
    'scheduled',
    'live',
    'completed',
    'cancelled'
    );

CREATE TYPE toss_decision AS ENUM (
    'bat',
    'bowl'
    );

CREATE TYPE batting_style AS ENUM (
    'right_hand_bat',
    'left_hand_bat'
    );

CREATE TYPE extra_type AS ENUM (
    'wide',
    'no_ball',
    'bye',
    'leg_bye'
    );

CREATE TYPE wicket_type AS ENUM (
    'bowled',
    'caught',
    'lbw',
    'run_out',
    'stumped',
    'retired_hurt'
    );

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name TEXT NOT NULL,
                       phone_no TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       archived_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE player_stats (
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                              runs BIGINT DEFAULT 0,
                              balls_faced BIGINT DEFAULT 0,
                              innings_batted BIGINT DEFAULT 0,
                              not_outs BIGINT DEFAULT 0,
                              fours BIGINT DEFAULT 0,
                              sixes BIGINT DEFAULT 0,
                              highest_score INT DEFAULT 0,
                              ducks BIGINT DEFAULT 0,
                              golden_ducks BIGINT DEFAULT 0,
                              fifties BIGINT DEFAULT 0,
                              hundreds BIGINT DEFAULT 0,
                              wickets BIGINT DEFAULT 0,
                              balls_bowled BIGINT DEFAULT 0,
                              runs_conceded BIGINT DEFAULT 0,
                              maiden_overs BIGINT DEFAULT 0,
                              wides BIGINT DEFAULT 0,
                              no_balls BIGINT DEFAULT 0,
                              innings_bowled BIGINT DEFAULT 0,
                              catches BIGINT DEFAULT 0,
                              run_outs BIGINT DEFAULT 0,
                              matches_played BIGINT DEFAULT 0,
                              matches_won BIGINT DEFAULT 0,
                              matches_lost BIGINT DEFAULT 0,
                              total_points BIGINT DEFAULT 0,
                              mvps BIGINT DEFAULT 0,
                              is_common_player BOOLEAN DEFAULT FALSE,
                              batting_style batting_style DEFAULT 'right_hand_bat',
                              created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              archived_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE teams (
                       team_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name VARCHAR(255) NOT NULL,
                       created_by UUID REFERENCES users(id) ON DELETE SET NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       archived_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE team_players (
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              team_id UUID NOT NULL REFERENCES teams(team_id) ON DELETE CASCADE,
                              player_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                              created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              UNIQUE(team_id, player_id)
);

CREATE TABLE matches (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         host_user_id UUID REFERENCES users(id) ON DELETE SET NULL,
                         team1_id UUID NOT NULL REFERENCES teams(team_id),
                         team2_id UUID NOT NULL REFERENCES teams(team_id),
                         venue VARCHAR(255),
                         overs INT NOT NULL CHECK (overs > 0),
                         players_per_team INT DEFAULT 11,
                         status match_status DEFAULT 'scheduled',
                         toss_winner_id UUID REFERENCES teams(team_id),
                         toss_decision toss_decision,
                         winner_team_id UUID REFERENCES teams(team_id),
                         man_of_match_id UUID REFERENCES users(id),
                         worst_player_id UUID REFERENCES users(id),
                         started_at TIMESTAMP WITH TIME ZONE,
                         ended_at TIMESTAMP WITH TIME ZONE,
                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         archived_at TIMESTAMP WITH TIME ZONE,
                         CHECK (team1_id <> team2_id)
);

CREATE TABLE match_players (
                               id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               match_id UUID NOT NULL REFERENCES matches(id) ON DELETE CASCADE,
                               player_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                               team_id UUID NOT NULL REFERENCES teams(team_id) ON DELETE CASCADE,
                               is_captain BOOLEAN DEFAULT FALSE,
                               created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                               UNIQUE(match_id, player_id)
);



CREATE TABLE innings (

                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

                         match_id UUID NOT NULL REFERENCES matches(id) ON DELETE CASCADE,

                         inning_number INT NOT NULL,

                         batting_team_id UUID NOT NULL REFERENCES teams(team_id),

                         bowling_team_id UUID NOT NULL REFERENCES teams(team_id),

                         total_runs INT DEFAULT 0,

                         wickets INT DEFAULT 0,

                         overs DECIMAL(4,1) DEFAULT 0.0,

                         extras INT DEFAULT 0,

                         target INT,

                         completed_overs INT DEFAULT 0,

                         balls_in_current_over INT DEFAULT 0,

                         created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

                         UNIQUE(match_id, inning_number)
);

CREATE TABLE deliveries (

                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

                            inning_id UUID NOT NULL REFERENCES innings(id) ON DELETE CASCADE,

                            over_number INT NOT NULL,

                            ball_number INT NOT NULL,

                            striker_id UUID REFERENCES users(id),

                            non_striker_id UUID REFERENCES users(id),

                            bowler_id UUID REFERENCES users(id),

                            runs_bat INT DEFAULT 0,

                            extras INT DEFAULT 0,

                            extra_type extra_type,

                            total_runs INT DEFAULT 0,

                            wicket BOOLEAN DEFAULT FALSE,

                            wicket_type wicket_type,

                            fielder_id UUID REFERENCES users(id),

                            player_out_id UUID REFERENCES users(id),

                            is_free_hit BOOLEAN DEFAULT FALSE,

                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

                            UNIQUE(inning_id, over_number, ball_number)
);

CREATE TABLE batting_scorecards (

                                    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

                                    inning_id UUID NOT NULL REFERENCES innings(id) ON DELETE CASCADE,

                                    player_id UUID NOT NULL REFERENCES users(id),

                                    runs INT DEFAULT 0,

                                    balls_faced INT DEFAULT 0,

                                    fours INT DEFAULT 0,

                                    sixes INT DEFAULT 0,

                                    strike_rate DECIMAL(6,2) DEFAULT 0,

                                    dismissal_type wicket_type,

                                    bowler_id UUID REFERENCES users(id),

                                    fielder_id UUID REFERENCES users(id),

                                    is_out BOOLEAN DEFAULT FALSE,

                                    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

                                    UNIQUE(inning_id, player_id)
);

CREATE TABLE bowling_scorecards (

                                    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

                                    inning_id UUID NOT NULL REFERENCES innings(id) ON DELETE CASCADE,

                                    player_id UUID NOT NULL REFERENCES users(id),

                                    overs DECIMAL(4,1) DEFAULT 0.0,

                                    maidens INT DEFAULT 0,

                                    runs_conceded INT DEFAULT 0,

                                    wickets INT DEFAULT 0,

                                    wides INT DEFAULT 0,

                                    no_balls INT DEFAULT 0,

                                    economy DECIMAL(6,2) DEFAULT 0,

                                    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

                                    UNIQUE(inning_id, player_id)
);

CREATE INDEX idx_users_phone_no
    ON users(phone_no);

CREATE INDEX idx_player_stats_user_id
    ON player_stats(user_id);

CREATE INDEX idx_team_players_team_id
    ON team_players(team_id);

CREATE INDEX idx_team_players_player_id
    ON team_players(player_id);

CREATE INDEX idx_matches_status
    ON matches(status);

CREATE INDEX idx_innings_match_id
    ON innings(match_id);

CREATE INDEX idx_deliveries_inning_id
    ON deliveries(inning_id);

CREATE INDEX idx_batting_scorecards_inning_id
    ON batting_scorecards(inning_id);

CREATE INDEX idx_bowling_scorecards_inning_id
    ON bowling_scorecards(inning_id);

ALTER TABLE deliveries
    ADD COLUMN is_legal_delivery BOOLEAN DEFAULT TRUE;

CREATE TABLE live_match_stats (
                                  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

                                  match_id UUID NOT NULL UNIQUE,
                                  innings_id UUID NOT NULL,

                                  batting_team_id UUID,
                                  bowling_team_id UUID,

                                  striker_id UUID,
                                  non_striker_id UUID,
                                  bowler_id UUID,

                                  current_score INTEGER DEFAULT 0,
                                  wickets INTEGER DEFAULT 0,
                                  legal_balls INTEGER DEFAULT 0,

                                  current_over INTEGER DEFAULT 0,

                                  required_runs INTEGER DEFAULT 0,

                                  last_updated TIMESTAMP DEFAULT NOW()
);

select * from matches