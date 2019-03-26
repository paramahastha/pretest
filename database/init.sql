--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.12
-- Dumped by pg_dump version 10.6 (Ubuntu 10.6-0ubuntu0.18.04.1)

-- Started on 2019-03-26 11:02:19 WIB

ALTER DATABASE pretest_db OWNER TO postgres;

\connect pretest_db

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 12393)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2179 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- TOC entry 1 (class 3079 OID 24576)
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- TOC entry 2180 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


--
-- TOC entry 187 (class 1259 OID 24590)
-- Name: ai_category; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ai_category
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ai_category OWNER TO postgres;

--
-- TOC entry 189 (class 1259 OID 24594)
-- Name: ai_image; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ai_image
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ai_image OWNER TO postgres;

--
-- TOC entry 188 (class 1259 OID 24592)
-- Name: ai_product; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ai_product
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ai_product OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 186 (class 1259 OID 24585)
-- Name: category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category (
    id integer DEFAULT nextval('public.ai_category'::regclass) NOT NULL,
    enable boolean NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.category OWNER TO postgres;

--
-- TOC entry 191 (class 1259 OID 24605)
-- Name: category_product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category_product (
    product_id integer NOT NULL,
    category_id integer NOT NULL
);


ALTER TABLE public.category_product OWNER TO postgres;

--
-- TOC entry 192 (class 1259 OID 24620)
-- Name: image; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.image (
    id integer DEFAULT nextval('public.ai_image'::regclass) NOT NULL,
    name character varying(64) NOT NULL,
    file character varying(255) NOT NULL,
    enable boolean NOT NULL
);


ALTER TABLE public.image OWNER TO postgres;

--
-- TOC entry 190 (class 1259 OID 24596)
-- Name: product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product (
    id integer DEFAULT nextval('public.ai_product'::regclass) NOT NULL,
    name character varying(64) NOT NULL,
    description character varying(255) NOT NULL,
    enable boolean NOT NULL
);


ALTER TABLE public.product OWNER TO postgres;

--
-- TOC entry 193 (class 1259 OID 24625)
-- Name: product_image; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_image (
    product_id integer NOT NULL,
    image_id integer NOT NULL
);


ALTER TABLE public.product_image OWNER TO postgres;

--
-- TOC entry 2164 (class 0 OID 24585)
-- Dependencies: 186
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category (id, enable, name) FROM stdin;
1	t	book
\.


--
-- TOC entry 2169 (class 0 OID 24605)
-- Dependencies: 191
-- Data for Name: category_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category_product (product_id, category_id) FROM stdin;
1	1
\.


--
-- TOC entry 2170 (class 0 OID 24620)
-- Dependencies: 192
-- Data for Name: image; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.image (id, name, file, enable) FROM stdin;
1	harry potter	https://images.ctfassets.net/bxd3o8b291gf/7cpPZFFu92oY46WWiiQcc2/909c7290404dfcda472cb330dc84c135/Harry-Potter-in-cupboard-Jim-Kay-RGB-636x800.jpg?w=1100&q=85	t
\.


--
-- TOC entry 2168 (class 0 OID 24596)
-- Dependencies: 190
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product (id, name, description, enable) FROM stdin;
1	harry potter	goblet of fire	t
\.


--
-- TOC entry 2171 (class 0 OID 24625)
-- Dependencies: 193
-- Data for Name: product_image; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_image (product_id, image_id) FROM stdin;
1	1
\.


--
-- TOC entry 2181 (class 0 OID 0)
-- Dependencies: 187
-- Name: ai_category; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ai_category', 1, true);


--
-- TOC entry 2182 (class 0 OID 0)
-- Dependencies: 189
-- Name: ai_image; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ai_image', 1, true);


--
-- TOC entry 2183 (class 0 OID 0)
-- Dependencies: 188
-- Name: ai_product; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ai_product', 1, true);


--
-- TOC entry 2030 (class 2606 OID 24604)
-- Name: category category_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_id_key UNIQUE (id);


--
-- TOC entry 2032 (class 2606 OID 24589)
-- Name: category category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- TOC entry 2038 (class 2606 OID 24619)
-- Name: category_product category_product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_product
    ADD CONSTRAINT category_product_pkey PRIMARY KEY (product_id, category_id);


--
-- TOC entry 2040 (class 2606 OID 24624)
-- Name: image image_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.image
    ADD CONSTRAINT image_pkey PRIMARY KEY (id);


--
-- TOC entry 2034 (class 2606 OID 24602)
-- Name: product product_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_id_key UNIQUE (id);


--
-- TOC entry 2042 (class 2606 OID 24629)
-- Name: product_image product_image_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_image
    ADD CONSTRAINT product_image_pkey PRIMARY KEY (product_id, image_id);


--
-- TOC entry 2036 (class 2606 OID 24600)
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- TOC entry 2043 (class 2606 OID 24608)
-- Name: category_product category_product_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_product
    ADD CONSTRAINT category_product_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.category(id);


--
-- TOC entry 2044 (class 2606 OID 24613)
-- Name: category_product category_product_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_product
    ADD CONSTRAINT category_product_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product(id);


--
-- TOC entry 2046 (class 2606 OID 24635)
-- Name: product_image product_image_image_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_image
    ADD CONSTRAINT product_image_image_id_fkey FOREIGN KEY (image_id) REFERENCES public.image(id);


--
-- TOC entry 2045 (class 2606 OID 24630)
-- Name: product_image product_image_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_image
    ADD CONSTRAINT product_image_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product(id);


-- Completed on 2019-03-26 11:02:19 WIB

--
-- PostgreSQL database dump complete
--

