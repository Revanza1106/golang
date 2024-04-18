CREATE TABLE produk
(
 id serial NOT NULL,
 Nama Produk character varying NOT NULL,
 Jenis character varying,
exp date,
 CONSTRAINT pk_produk PRIMARY KEY (id )
)
WITH (
 OIDS=FALSE
);
ALTER TABLE produk
 OWNER TO postgres;