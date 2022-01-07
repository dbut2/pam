CREATE TABLE IF NOT EXISTS entry (
    id varchar(36) NOT NULL
);

CREATE TABLE IF NOT EXISTS entry_history (
    id       varchar(36) NOT NULL,
    metadata varchar(36) NOT NULL,
    date     date,
    entry    text
);

CREATE TABLE IF NOT EXISTS user (
    id  varchar(36) NOT NULL,
    gid varchar(36)
);

CREATE TABLE IF NOT EXISTS metadata (
    id        varchar(36)  NOT NULL,
    author    varchar(36)  NOT NULL,
    timestamp timestamp(6) NOT NULL
);

ALTER TABLE entry
    ADD PRIMARY KEY (id);

ALTER TABLE entry_history
    ADD PRIMARY KEY (metadata),
    ADD KEY id (id),
    ADD KEY date (date);

ALTER TABLE user
    ADD PRIMARY KEY (id),
    ADd KEY id (id),
    ADD KEY gid (gid);

ALTER TABLE metadata
    ADD PRIMARY KEY (id),
    ADD KEY id (id),
    ADD KEY author (author);

ALTER TABLE entry_history
    ADD CONSTRAINT entry_history_entry FOREIGN KEY (id) REFERENCES entry (id) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT entry_history_metadata FOREIGN KEY (metadata) REFERENCES metadata (id) ON DELETE CASCADE ON UPDATE CASCADE;
