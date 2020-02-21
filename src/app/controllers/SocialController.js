import Social from '../models/Social';

class SocialController {
  async store(req) {
    Social.create(req);
  }
}

export default new SocialController();
