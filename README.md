# CoderSchool Golang Course - Week 1.3 Assignment: URL

1. **Submitted by: Lam Tran**
2. **Time spent: 8**

## Set of User Stories

### Required

- [x] Manipulate YAML config file. Where the redirection list peristently stored.
- [x] Implement append to the list: `urlshorten configure -a dogs -u www.dogs.com`
- [x] Implement remove from the list: `urlshorten -d dogs`
- [x] List redirections: `urlshorten -l`
- [x] Run HTTP server on a given port: `urlshorten run -p 8080`
- [x] Prints usage info: `urlshorten -h`

### Bonus

- [ ] Track number of times each redirection is used. When the user uses `urlshorten -l`, the user should see redirections ranked by how many times they have been used.
- [ ] Provide a default shortening, if no example is given. For example, if `dogs` is not given, generate a random alphanumeric string of length 8.
- [ ] Build a Handler that doesn't read from a map but instead reads from a database. Whether you use BoltDB, SQL, or something else is entirely up to you.
