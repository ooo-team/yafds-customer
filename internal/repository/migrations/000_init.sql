-- public.customers definition

-- Drop table

DROP TABLE IF EXISTS public.customers;
go;

CREATE TABLE public.customers (
	id int4 NOT NULL,
	phone varchar(20) NULL,
	email varchar NULL,
	address varchar NULL
);
CREATE UNIQUE INDEX xpkcustomers ON public.customers USING btree (id);
CREATE UNIQUE INDEX xie1customers ON public.customers USING btree (phone);
CREATE UNIQUE INDEX xie2customers ON public.customers USING btree (email);


-- public.h_customers definition

-- Drop table

DROP TABLE IF EXISTS public.h_customers;
go;

CREATE TABLE public.h_customers (
	customer_id int4 NOT NULL,
	created_at varchar NOT NULL,
	modified_at varchar NULL
);
CREATE UNIQUE INDEX xpkhcustomers ON public.h_customers USING btree (customer_id, modified_at);