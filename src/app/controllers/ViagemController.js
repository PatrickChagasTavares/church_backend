import ChildrenController from './ChildrenController';
import DoorToDoorController from './DoorToDoorController';
import SocialController from './SocialController';

class ViagemController {
  async store(req, res) {
    try {
      const { children, DoorToDoor, Social } = req.body;

      if (children && children.length > 0) {
        await children.map(item => ChildrenController.store(item));
      }

      if (DoorToDoor && DoorToDoor.length > 0) {
        await DoorToDoor.map(item => DoorToDoorController.store(item));
      }
      if (Social && Social.length > 0) {
        await Social.map(item => SocialController.store(item));
      }

      return res.json({ status: true });
    } catch (error) {
      console.log(error);
      return res.json({ status: 'error' });
    }
  }
}

export default new ViagemController();
