import DoorToDoor from '../models/DoorToDoor';

class DoorToDoorController {
  async store(req) {
    DoorToDoor.create(req);
  }
}

export default new DoorToDoorController();
