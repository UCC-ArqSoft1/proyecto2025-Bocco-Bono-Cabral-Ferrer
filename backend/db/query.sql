CREATE DATABASE db_gym;
insert into users values(2,"Hulk","Hogan", "user@user.com","04f8996da763b7a969b1028ee3007569eaf3a635486ddab211d512c85b9df8fb","19/19/1919","minita",2);
insert into users values(1,"nic","Bon", "admn@admn.com","0d335a3bea76dac4e3926d91c52d5bdd716bac2b16db8caf3fb6b7a58cbd92a7","19/19/1919","minita",1);
insert into user_types values(1,"admn");
insert into user_types values(2,"usuario");
INSERT INTO activities (
    name, description, capacity, category, profesor, image_url
) VALUES 
(
    'Yoga para principiantes',
    'Clase de yoga enfocada en la respiración y estiramientos básicos.',
    20,
    'Yoga',
    'Laura Martínez',
    'images/yoga.jpg'
),
(
    'CrossFit Avanzado',
    'Entrenamiento funcional de alta intensidad para usuarios avanzados.',
    15,
    'CrossFit',
    'Carlos Gómez',
    'images/crosfit.jpg'
),
(
    'Spinning',
    'Clase de ciclismo indoor con música motivadora.',
    25,
    'Cardio',
    'Ana Rodríguez',
    'images/spinning.jpg'
);
INSERT INTO activity_schedules(
    activity_id, day, start_time, end_time
) VALUES 
-- Yoga: Lunes y Miércoles de 08:00 a 09:00
(1, 'Monday', '08:00', '09:00'),
(1, 'Wednesday', '08:00', '09:00'),

-- CrossFit: Martes y Jueves de 18:00 a 19:00
(2, 'Tuesday', '18:00', '19:00'),
(2, 'Thursday', '18:00', '19:00'),

-- Spinning: Viernes de 17:00 a 18:00 y Sábado de 10:00 a 11:00
(3, 'Friday', '17:00', '18:00'),
(3, 'Saturday', '10:00', '11:00');

