import { UserService } from "./user.service";

export class ConfigService {
  private userService: UserService;

  constructor() {
    this.userService = new UserService();
  }

  getUserConfig(userId: string) {
    const user = this.userService.getUser(userId);
    return { theme: "dark", language: "en", user };
  }

  updateConfig(userId: string, config: any) {
    return true;
  }
}
