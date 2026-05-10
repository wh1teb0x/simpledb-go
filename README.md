# simpledb-go

A Go implementation of [SimpleDB](https://cs.bc.edu/~sciore/simpledb/), the educational RDBMS from
*Database Design and Implementation* by Edward Sciore (Springer).

This is a from-scratch port written in idiomatic Go for self-study — not a line-for-line translation
of the Java reference code.

## Roadmap

Built bottom-up, following the textbook:

- [ ] `file` — block-based disk I/O (`BlockId`, `Page`, `FileMgr`)
- [ ] `log` — write-ahead log
- [ ] `buffer` — buffer pool
- [ ] `tx` — transactions, locking, recovery
- [ ] `record` — records, schemas, table scans
- [ ] `metadata` — system catalog
- [ ] `query` / `parse` / `plan` — SQL execution
- [ ] `index` / `materialize` / `opt` — indexing & query optimization

## Reference

- Java reference: <https://github.com/LutherCS/sciore-simpledb-pub>
- Textbook: Sciore, *Database Design and Implementation*, 2nd ed., Springer.
