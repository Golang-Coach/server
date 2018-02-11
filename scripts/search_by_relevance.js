db.repositories.aggregate(
   [
     { $match: { $text: { $search: "react" } } },
     { $sort: { score: { $meta: "textScore" } } },
     { $skip : 2 },
     { $limit: 20 }
   ]
)