CREATE TYPE state_type AS ENUM (
    'dead',
    'exited',
    'waiting',
    'running',
    'failed'
);

CREATE TABLE service (
    id           integer     PRIMARY KEY GENERATED ALWAYS as IDENTITY,
    name         varchar     NOT NULL, 
    object_path  varchar     NOT NULL,
    state        state_type  NOT NULL
);

CREATE UNIQUE INDEX udx__service__name ON service(name);

CREATE TABLE subscription (
    id           integer     PRIMARY KEY GENERATED ALWAYS as IDENTITY,
    service_id   integer     REFERENCES  service(id)
);

CREATE UNIQUE INDEX udx__subscription__service_id ON subscription(service_id);
