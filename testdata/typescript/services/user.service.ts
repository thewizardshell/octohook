export class UserService {
  getUser(id: string) {
    return { id, name: "John Doe" };
  }

  createUser(name: string) {
    return { id: "123", name };
  }

  deleteUser(id: string) {
    return true;
  }
}
