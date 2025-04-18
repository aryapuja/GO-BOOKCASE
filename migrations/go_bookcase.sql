PGDMP  	    9                }            go_bookcase    16.8    16.8     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16418    go_bookcase    DATABASE     q   CREATE DATABASE go_bookcase WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en-US';
    DROP DATABASE go_bookcase;
                postgres    false            �            1259    16430    books    TABLE     H  CREATE TABLE public.books (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    image_url character varying(255),
    release_year integer,
    price integer,
    total_page integer,
    thickness character varying(50),
    category_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by character varying(255),
    modified_at timestamp without time zone,
    modified_by character varying(255),
    CONSTRAINT books_release_year_check CHECK (((release_year >= 1980) AND (release_year <= 2024)))
);
    DROP TABLE public.books;
       public         heap    postgres    false            �            1259    16429    books_id_seq    SEQUENCE     �   CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.books_id_seq;
       public          postgres    false    218            �           0    0    books_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;
          public          postgres    false    217            �            1259    16420 
   categories    TABLE     (  CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by character varying(255),
    modified_at timestamp without time zone,
    modified_by character varying(255)
);
    DROP TABLE public.categories;
       public         heap    postgres    false            �            1259    16419    categories_id_seq    SEQUENCE     �   CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.categories_id_seq;
       public          postgres    false    216            �           0    0    categories_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;
          public          postgres    false    215            �            1259    16446    users    TABLE     U  CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by character varying(255),
    modified_at timestamp without time zone,
    modified_by character varying(255)
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    16445    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    220            �           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    219            &           2604    16433    books id    DEFAULT     d   ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);
 7   ALTER TABLE public.books ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    217    218            $           2604    16423    categories id    DEFAULT     n   ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);
 <   ALTER TABLE public.categories ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    216    216            (           2604    16449    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    220    219    220            �          0    16430    books 
   TABLE DATA           �   COPY public.books (id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by) FROM stdin;
    public          postgres    false    218   y       �          0    16420 
   categories 
   TABLE DATA           `   COPY public.categories (id, name, created_at, created_by, modified_at, modified_by) FROM stdin;
    public          postgres    false    216   �#       �          0    16446    users 
   TABLE DATA           i   COPY public.users (id, username, password, created_at, created_by, modified_at, modified_by) FROM stdin;
    public          postgres    false    220   $       �           0    0    books_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.books_id_seq', 27, true);
          public          postgres    false    217            �           0    0    categories_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.categories_id_seq', 7, true);
          public          postgres    false    215            �           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 2, true);
          public          postgres    false    219            .           2606    16439    books books_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.books DROP CONSTRAINT books_pkey;
       public            postgres    false    218            ,           2606    16428    categories categories_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.categories DROP CONSTRAINT categories_pkey;
       public            postgres    false    216            0           2606    16454    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    220            2           2606    16456    users users_username_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_key;
       public            postgres    false    220            3           2606    16440    books books_category_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id);
 F   ALTER TABLE ONLY public.books DROP CONSTRAINT books_category_id_fkey;
       public          postgres    false    4652    218    216            �     x���M��8���_�㮴�|B���L}Zͥ��8q�?zĿ߲���X	�i �v��[U�3!�@Pt�J�{N^�OVRä�������H��`e]�jAA��Zp�`%7V�ɈDA��0p$��a;(I��t$�0�a:K�Y���<��)����|���)Y؊�� ��В�F��;S܀� GPp�+YԒ�L.���
���;Gv�	�, a�i��i2��(Ɍ.1P��aɺD��^	�
A��a� �*aBr��Oc�d'Τ��9](�>�W�M�δ�C[!�J'�%m�DU#%-P6�Y��C��a�L�jr�&z*ȟ`�%��ʅ\������5ʉa��P��5SA��5�@s�����g�'X��ZU8�4�|&�Q0�>��L1M-�1Uړ��e%؁8��5jV���%v�
@`ԕZ�`D��g5~�)"����ur�\�`ع��Jc�co52j�R�>[���B�q��d^",�'~'&oLȺ�����.���e���j���A�c�	��UX��X�0���ǁ���÷�/,�o`�uu1��҆U�kn	�.A� %�[�A+ ��#���RhPL�7���=�j�&$u.��<7!�ޘ7p(QI_�ej��n�vx��ڃ��K����e��vi�{}t:H�?��X/��B�͇=v��n��n�1�v�8{�-C�%��:����Y�z���Ӎ����4H��A�V��G_�d�s���]�_���|o5���YL-��09C���vn��y4���m��?��u�Ҙ)�0�����GZE����J!^���0�q��YB�ݶ��Z�^8ř�]w��_�M�[���n��).y�5��=y�0W4�����Nqg�S�0p����w��e���d�j�w�*[�a`ҹрb����n�[I�ʒ�]����8�I���m���'�'K���'��΍�c>bD�g�<'�����d�(O�8˦��&�xo�?�c���IEI�w���L�1>��8��y=!�gQ0��4���s����/96C       �   v   x�3�t�L.����4202�50�54Q04�25�2�Գ�44���,�,.I���".#N��<]Rus'g��%��Ä�#��$���h��N���E���1�IM�����O'^S� ijN{      �   x   x�3�LL����T1JR14R��r2��(���r)���u�4�	�
pN5�q446�4�������p��)�5���4202�50�54Q04�22�2��360�42�,�,.I���"�=... ���     