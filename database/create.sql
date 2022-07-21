-- Startrail database table description.

-- User table for general information
CREATE TABLE users (
  nickname        varchar         primary key,
  descr           text,
  email           varchar,
  password        varchar(100)    not null,
  is_admin        boolean         default false
);

-- This table implements the friend system. 
-- A is B's friend iff (A, B, true) and (B, A, true) exist in the database.
-- If (A, B, true) and (B, A, false), then A sent a friend request to B,
-- and B must accept it in order to become friends.
CREATE TABLE befriends (
  nickname1       varchar         not null,
  nickname2       varchar         not null,
  accepted        boolean         default false,

  CONSTRAINT PK_bf PRIMARY KEY(nickname1, nickname2),
  CONSTRAINT FK_bf_n1 FOREIGN KEY(nickname1) REFERENCES users(nickname) ON DELETE CASCADE,
  CONSTRAINT FK_bf_n2 FOREIGN KEY(nickname2) REFERENCES users(nickname) ON DELETE CASCADE
);

-- Stat table, which contains the 
CREATE TABLE stats (
  tstamp          date            not null,
  steps           integer         not null,
  kcal            decimal         not null,
  distance        decimal         not null,
  nickname        varchar         not null,

  CONSTRAINT PK_stats PRIMARY KEY(nickname, tstamp),
  CONSTRAINT FK_stats_nickname FOREIGN KEY(nickname) REFERENCES users(nickname) ON DELETE CASCADE
);

-- Mission types are identified by the following numbers:
-- 1 -> steps
-- 2 -> kcal
-- 3 -> distance
CREATE TABLE missions (
  id              serial,
  tstamp          timestamp        not null,
  mtype           integer         not null,

  -- val sets the value that must be reached in order to complete the mission
  val             integer         not null,
  nickname        varchar         not null,

  CONSTRAINT PK_ms PRIMARY KEY(nickname, id),
  CONSTRAINT FK_ms_nickname FOREIGN KEY(nickname) REFERENCES users(nickname) ON DELETE CASCADE
);

-- Events contain the stats acquired by any user in a given period of time.
-- They are organized by the app's admin
CREATE TABLE events (
  id              serial,
  date_start      date            not null,
  date_end        date            not null,
  steps           integer         not null,
  kcal            decimal         not null,
  distance        decimal         not null,
  nickname        varchar         not null,

  CONSTRAINT PK_ev PRIMARY KEY(nickname, id),
  CONSTRAINT FK_ev_nickname FOREIGN KEY(nickname) REFERENCES users(nickname) ON DELETE CASCADE
);

CREATE TABLE achievements (
  id              serial          primary key,
  tstamp          timestamp       not null
);

CREATE TABLE user_achievements (
  id              integer         not null,
  nickname        varchar         not null,
  tstamp          timestamp       not null,

  CONSTRAINT PK_ua PRIMARY KEY(nickname, id),
  CONSTRAINT FK_ua_id FOREIGN KEY(id) REFERENCES achievements(id) ON DELETE CASCADE,
  CONSTRAINT FK_ua_nickname FOREIGN KEY(nickname) REFERENCES users(nickname) ON DELETE CASCADE
);