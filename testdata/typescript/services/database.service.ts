import { AuthService } from "./auth.service";
import { ConfigService } from "./config.service";

export class DatabaseService {
  private authService: AuthService;
  private configService: ConfigService;

  constructor() {
    this.authService = new AuthService();
    this.configService = new ConfigService();
  }

  connect() {
    return true;
  }

  saveUserData(userId: string, data: any) {
    const config = this.configService.getUserConfig(userId);
    return { saved: true, config };
  }
}
