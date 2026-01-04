insert into faculties (name) values
    ('Engineering'),
    ('Arts'),
    ('Science'),
    ('Business');

insert into groups (name, faculty_id) values
    ('ENG101', 1),
    ('CSI201', 2),
    ('WCS301', 3),
    ('BUS401', 4);

insert into students (first_name, last_name, sex, birth_date, group_id) values
    ('Alice', 'Smith', 'f', '2000-05-15', 1),
    ('Bob', 'Johnson', 'm', '1999-08-22', 2),
    ('Catherine', 'Williams', 'f', '2001-12-03', 3),
    ('David', 'Brown', 'm', '2000-11-30', 4),
    ('Eva', 'Davis', 'f', '1998-07-19', 1),
    ('Frank', 'Miller', 'm', '1997-03-25', 2),
    ('Grace', 'Wilson', 'f', '2002-01-10', 3),
    ('Henry', 'Moore', 'm', '1999-09-14', 4);

insert into schedule (group_id, faculty_id, subject, day_of_week, start_time, end_time) values
    (1, 1, 'Mathematics', 'mon', '09:00', '10:30'),
    (2, 2, 'Literature', 'tue', '11:00', '12:30'),
    (3, 3, 'Physics', 'wed', '13:00', '14:30'),
    (4, 4, 'Economics', 'thu', '15:00', '16:30'),
    (1, 1, 'Chemistry', 'fri', '10:00', '11:30'),
    (2, 2, 'Mathematics', 'mon', '12:00', '13:30'),
    (3, 3, 'Literature', 'tue', '14:00', '15:30'),
    (4, 4, 'Physics', 'wed', '16:00', '17:30');