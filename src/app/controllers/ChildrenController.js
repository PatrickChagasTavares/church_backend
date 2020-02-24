import Children from '../models/Children';

class ChildrenController {
  async store(req) {
    Children.create(req);
  }
}

export default new ChildrenController();
