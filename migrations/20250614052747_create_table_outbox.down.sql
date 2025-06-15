BEGIN;

DROP TABLE IF EXISTS outbox;

DROP TYPE outbox_status;

DROP TYPE outbox_event;

COMMIT;
