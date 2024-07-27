package pg

func NewRepo(db *NichirinPg) NichirinRepo {
	return NichirinRepo{
		Pg: db,
		Db: db.Db,
	}
}
