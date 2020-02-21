import Children from '../models/Children';

class ChildrenController {
  async store(req) {
    console.log(req);
    Children.create(req);
  }
}

export default new ChildrenController();
