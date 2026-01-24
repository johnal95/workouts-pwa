-- +goose Up
-- +goose StatementBegin
INSERT INTO exercises (id, type)
VALUES
    ('019bed7e-0ef7-7401-85f4-000000000001', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000002', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000003', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000004', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000005', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000006', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000007', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000008', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000009', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000010', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000011', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000012', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000013', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000014', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000015', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000016', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000017', 'REPS'),
    ('019bed7e-0ef7-7401-85f4-000000000018', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000019', 'REPS'),
    ('019bed7e-0ef7-7401-85f4-000000000020', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000021', 'DURATION'),
    ('019bed7e-0ef7-7401-85f4-000000000022', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000023', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000024', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000025', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000026', 'REPS'),
    ('019bed7e-0ef7-7401-85f4-000000000027', 'REPS_WEIGHT'),
    ('019bed7e-0ef7-7401-85f4-000000000028', 'DURATION');

INSERT INTO exercise_translations (exercise_id, locale, name)
VALUES
    ('019bed7e-0ef7-7401-85f4-000000000001', 'en-US', 'Barbell Row'),
    ('019bed7e-0ef7-7401-85f4-000000000002', 'en-US', 'Cable Close Grip Seated Row'),
    ('019bed7e-0ef7-7401-85f4-000000000003', 'en-US', 'Cable Crossover'),
    ('019bed7e-0ef7-7401-85f4-000000000004', 'en-US', 'Deadlift'),
    ('019bed7e-0ef7-7401-85f4-000000000005', 'en-US', 'Decline Barbell Press'),
    ('019bed7e-0ef7-7401-85f4-000000000006', 'en-US', 'Decline Dumbbell Press'),
    ('019bed7e-0ef7-7401-85f4-000000000007', 'en-US', 'Dumbbell Lateral Raise'),
    ('019bed7e-0ef7-7401-85f4-000000000008', 'en-US', 'Dumbbell Lunges'),
    ('019bed7e-0ef7-7401-85f4-000000000009', 'en-US', 'Flat Barbell Press'),
    ('019bed7e-0ef7-7401-85f4-000000000010', 'en-US', 'Flat Dumbbell Press'),
    ('019bed7e-0ef7-7401-85f4-000000000011', 'en-US', 'Incline Barbell Press'),
    ('019bed7e-0ef7-7401-85f4-000000000012', 'en-US', 'Incline Dumbbell Press'),
    ('019bed7e-0ef7-7401-85f4-000000000013', 'en-US', 'Leg Extension'),
    ('019bed7e-0ef7-7401-85f4-000000000014', 'en-US', 'Leg Press'),
    ('019bed7e-0ef7-7401-85f4-000000000015', 'en-US', 'Military Press'),
    ('019bed7e-0ef7-7401-85f4-000000000016', 'en-US', 'Overhand Lat Pulldown'),
    ('019bed7e-0ef7-7401-85f4-000000000017', 'en-US', 'Overhand Pull-Up'),
    ('019bed7e-0ef7-7401-85f4-000000000018', 'en-US', 'Parallel Grip Lat Pulldown'),
    ('019bed7e-0ef7-7401-85f4-000000000019', 'en-US', 'Parallel Grip Pull-Up'),
    ('019bed7e-0ef7-7401-85f4-000000000020', 'en-US', 'Pec Deck Fly'),
    ('019bed7e-0ef7-7401-85f4-000000000021', 'en-US', 'Plank'),
    ('019bed7e-0ef7-7401-85f4-000000000022', 'en-US', 'Seated Leg Curl'),
    ('019bed7e-0ef7-7401-85f4-000000000023', 'en-US', 'Seated Reverse Dumbbell Fly'),
    ('019bed7e-0ef7-7401-85f4-000000000024', 'en-US', 'Squat'),
    ('019bed7e-0ef7-7401-85f4-000000000025', 'en-US', 'Underhand Lat Pulldown'),
    ('019bed7e-0ef7-7401-85f4-000000000026', 'en-US', 'Underhand Pull-Up'),
    ('019bed7e-0ef7-7401-85f4-000000000027', 'en-US', 'Upright Row'),
    ('019bed7e-0ef7-7401-85f4-000000000028', 'en-US', 'Wall Sit');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM exercises;
-- +goose StatementEnd
