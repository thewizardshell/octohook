import { UserService } from "./user.service";

export class AuthService {
  private userService: UserService;

  constructor() {
    this.userService = new UserService();
  }

  login(username: string, password: string) {
    const user = this.userService.getUser(username);
    return { token: "abc123", user };
  }

  logout(userId: string) {
    return true;
  }
}
