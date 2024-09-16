-- +goose Up
INSERT INTO employee (id, username, first_name, last_name)
VALUES ('7189d52d-bfcb-415b-9c7a-e5c16d5cae61', 'ivanov', 'Stepan', 'Ivanov'),
  ('e430b712-acce-444a-977e-dc109e952420', 'petrov', 'Alexey', 'Petrov'),
  ('ef958565-260c-452f-a28c-6c325366f3c9', 'sidorov', 'Alexandr', 'Sidorov'),
  ('1c1ceef4-fd69-4f63-ac4b-f1cdf6568197', 'kuznetsov', 'Viktor', 'Kuznetsov'),
  ('92a911cc-4061-40bb-bd7a-48f7d7efe330', 'sokolov', 'Eugen', 'Sokolov'),
  ('9479f3ff-bb18-4352-9b3d-196df264a74a', 'popov', 'Vladimir', 'Popov'),
  ('3ef7b563-4196-4b5d-b4c3-53c1b5737c38', 'sobolev', 'Anton', 'Sobolev'),
  ('a64cd128-bb67-4caa-83ba-d40e16173f72', 'klimov', 'Alexandr', 'Klimov');

-- +goose Down
DELETE FROM employee WHERE username IN ('ivanov', 'petrov', 'sidorov', 'kuznetsov', 'sokolov', 'popov', 'sobolev', 'klimov');