create table personalities(
    id serial primary key,
    name varchar,
    history varchar
);

INSERT INTO personalities(name, history) VALUES
('Deodato Petit Wertheimer', 'Deodato Petit Wertheimer was a Brazilian doctor and politician, his first years of life were in São Paulo, but he soon moved to Nova Friburgo in the State of Rio de Janeiro and at the age of 11 he entered the Jesuits Anchieta College.'),
('Carmela Dutra', 'Carmela Teles Leite Dutra was the first lady of Brazil, from January 31, 1946 until her death, having been the wife of Eurico Gaspar Dutra, 16th President of Brazil. She was affectionately called Dona Santinha, due to her strong religiosity, making her husband open a little chapel in the Guanabara Palace.');
