CREATE TABLE account(
    id     varchar(36) NOT NULL,
    person varchar(36) NOT NULL
);

CREATE TABLE activity (
    id        varchar(36) NOT NULL,
    metadata  varchar(36) NOT NULL,
    startTime date        NOT NULL,
    endTime   date        NOT NULL
);

CREATE TABLE celebration (
    id       varchar(36) NOT NULL,
    metadata varchar(36) NOT NULL,
    time     date        NOT NULL
);

CREATE TABLE metadata (
    id             varchar(36) NOT NULL,
    owner          varchar(36) NOT NULL,
    title          text        NOT NULL,
    description    text        NOT NULL,
    createdAt      date        NOT NULL,
    lastModifiedAt date        NOT NULL
);

CREATE TABLE note (
    id       varchar(36) NOT NULL,
    metadata varchar(36) NOT NULL,
    content  text        NOT NULL
);

CREATE TABLE person (
    id         varchar(36) NOT NULL,
    metadata   varchar(36) NOT NULL,
    firstname  text        NOT NULL,
    middlename text        NOT NULL,
    lastname   text        NOT NULL,
    birthday   varchar(36) NOT NULL
);

ALTER TABLE account
    ADD PRIMARY KEY (id),
    ADD KEY id (id),
    ADD KEY metadata (metadata);

ALTER TABLE activity
    ADD PRIMARY KEY (id),
    ADD KEY id (id),
    ADD KEY metadata (metadata);

ALTER TABLE celebration
    ADD PRIMARY KEY (id),
    ADD KEY id (id),
    ADD KEY metadata (metadata);

ALTER TABLE metadata
    ADD PRIMARY KEY (id),
    ADD KEY id (id),
    ADD KEY owner (owner);

ALTER TABLE note
    ADD PRIMARY KEY (id),
    ADD KEY id (id),
    ADD KEY metadata (metadata);

ALTER TABLE person
    ADD PRIMARY KEY (id),
    ADD KEY id (id),
    ADD KEY metadata (metadata),
    ADD KEY birthday (birthday);

ALTER TABLE account
    ADD CONSTRAINT account_person FOREIGN KEY (person) REFERENCES person (id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE activity
    ADD CONSTRAINT activity_metadata FOREIGN KEY (metadata) REFERENCES metadata (id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE celebration
    ADD CONSTRAINT celebration_metadata FOREIGN KEY (metadata) REFERENCES metadata (id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE metadata
    ADD CONSTRAINT metadata_metadata FOREIGN KEY (owner) REFERENCES account (id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE note
    ADD CONSTRAINT note_metadata FOREIGN KEY (metadata) REFERENCES metadata (id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE person
    ADD CONSTRAINT person_metadata FOREIGN KEY (metadata) REFERENCES metadata (id) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT person_birthday FOREIGN KEY (birthday) REFERENCES celebration (id) ON DELETE CASCADE ON UPDATE CASCADE;
