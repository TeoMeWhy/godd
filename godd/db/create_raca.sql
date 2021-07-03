drop table if exists tb_raca;
create table tb_raca (

    raca VARCHAR(50),
    forca INT,
    destreza INT,
    sabedoria INT,
    inteligencia INT,
    constituicao INT,
    carisma INT

);

INSERT INTO tb_modifier VALUES ("humano",1,1,1,1,1,1);
INSERT INTO tb_modifier VALUES ("anao",2,0,1,0,2,0);
INSERT INTO tb_modifier VALUES ("elfo",0,2,1,1,0,0);
INSERT INTO tb_modifier VALUES ("halfling",0,1,0,0,0,1);

select * from tb_modifier;