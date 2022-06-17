-- Fill client
insert into client(email) values("test1@test.ru");
insert into client(email) values("test2@test.ru");
insert into client(email) values("test3@test.ru");
insert into client(email) values("test4@test.ru");
-- Status:
-- 0 - New
-- 1 - SUCCESS don't change
-- 2 - Fail
-- 3 - Error don't change
-- Fill payment
---- Something id_client
insert into payment(id_client, amount, currency, time_created, time_changed, status_pay)
values(0, 10, "$", current_timestamp, current_timestamp, 0);
insert into payment(id_client, amount, currency, time_created, time_changed, status_pay)
values(0, 100, "$", current_timestamp, current_timestamp, 0);
insert into payment(id_client, amount, currency, time_created, time_changed, status_pay)
values(1, 20, "$", current_timestamp, current_timestamp, 2);
-- Fill don't change
insert into payment(id_client, amount, currency, time_created, time_changed, status_pay)
values(2, 20, "$", current_timestamp, current_timestamp, 1);
insert into payment(id_client, amount, currency, time_created, time_changed, status_pay)
values(3, 100, "$", current_timestamp, current_timestamp, 3);
insert into payment(id_client, amount, currency, time_created, time_changed, status_pay)
values(3, 50, "$", current_timestamp, current_timestamp, 1);
insert into payment(id_client, amount, currency, time_created, time_changed, status_pay)
values(3, 51, "$", current_timestamp, current_timestamp, 1);
