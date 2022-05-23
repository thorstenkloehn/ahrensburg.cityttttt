--
-- PostgreSQL database dump
--

-- Dumped from database version 12.11
-- Dumped by pg_dump version 12.11

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: postgis; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;


--
-- Name: EXTENSION postgis; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION postgis IS 'PostGIS geometry, geography, and raster spatial types and functions';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: Bauliche_Anlagen; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Bauliche_Anlagen" (
    id integer NOT NULL,
    name character varying(255),
    "popupContent" text,
    "Foto" character varying(255),
    geom public.geometry(Point)
);


ALTER TABLE public."Bauliche_Anlagen" OWNER TO postgres;

--
-- Name: Bauliche_Anlagen_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Bauliche_Anlagen_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Bauliche_Anlagen_id_seq" OWNER TO postgres;

--
-- Name: Bauliche_Anlagen_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Bauliche_Anlagen_id_seq" OWNED BY public."Bauliche_Anlagen".id;


--
-- Name: Bauliche_Anlagen id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Bauliche_Anlagen" ALTER COLUMN id SET DEFAULT nextval('public."Bauliche_Anlagen_id_seq"'::regclass);


--
-- Data for Name: Bauliche_Anlagen; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Bauliche_Anlagen" (id, name, "popupContent", "Foto", geom) FROM stdin;
\.


--
-- Data for Name: spatial_ref_sys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.spatial_ref_sys (srid, auth_name, auth_srid, srtext, proj4text) FROM stdin;
\.


--
-- Name: Bauliche_Anlagen_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Bauliche_Anlagen_id_seq"', 1, false);


--
-- Name: Bauliche_Anlagen Bauliche_Anlagen_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Bauliche_Anlagen"
    ADD CONSTRAINT "Bauliche_Anlagen_pkey" PRIMARY KEY (id);


--
-- Name: sidx_Bauliche_Anlagen_geom; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "sidx_Bauliche_Anlagen_geom" ON public."Bauliche_Anlagen" USING gist (geom);


--
-- PostgreSQL database dump complete
--

