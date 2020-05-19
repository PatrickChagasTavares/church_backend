import moment from 'moment';
import ChildrenController from './ChildrenController';
import DoorToDoorController from './DoorToDoorController';
import SocialController from './SocialController';

class ViagemController {
  async store(req, res) {
    const status = { children: false, DoorToDoor: false, Social: false };
    try {
      const { children, DoorToDoor, Social } = req.body;

      if (children && children.length > 0) {
        await children.map(item => {
          status.children = true;
          const formtData = {
            date: moment(item.data).format('YYYY-MM-DD'),
            total: item.total || 0,
            note: item.note || '',
          };
          return ChildrenController.store(formtData);
        });
      }

      if (DoorToDoor && DoorToDoor.length > 0) {
        await DoorToDoor.map(item => {
          status.DoorToDoor = true;
          const formtData = {
            date: moment(item.data).format('YYYY-MM-DD'),
            name_tribe: item.nameTribe,
            name_person: item.namePerson,
            address: item.address,
            type: item.type,
            age: item.age,
            bible: item.bible,
            evangelical: item.evangelical,
            contact: item.contact,
            frequents_church: item.frequentsChurch,
            cult_home: item.cultHome,
            study_bible: item.studyBible,
            study_confimation: item.studyConfimation,
            prayer_request: item.prayerRequest,
            reconciled: item.reconciled,
            visit: item.visit,
            accept_christ: item.acceptChrist,
            note: item.note || '',
          };

          return DoorToDoorController.store(formtData);
        });
      }

      if (Social && Social.length > 0) {
        await Social.map(item => {
          status.Social = true;
          const formtData = {
            date: moment(item.data).format('YYYY-MM-DD'),
            name_tribe: item.nameTribe,
            name_person: item.namePerson,
            address: item.address,
            type: item.type,
            age: item.age,
            bible: item.bible,
            evangelical: item.evangelical,
            contact: item.contact,
            frequents_church: item.frequentsChurch,
            cult_home: item.cultHome,
            study_bible: item.studyBible,
            study_confimation: item.studyConfimation,
            prayer_request: item.prayerRequest,
            reconciled: item.reconciled,
            visit: item.visit,
            accept_christ: item.acceptChrist,
            medical: item.medical,
            optician: item.optician,
            pressure: item.pressure,
            glucose: item.glucose,
            aesthetics: item.aesthetics,
            cutting_hair: item.cuttingHair,
            hairstyle: item.hairstyle,
            Nail: item.Nail,
            Eyebrow: item.Eyebrow,
            note: item.note,
          };
          return SocialController.store(formtData);
        });
      }

      if (
        status.DoorToDoor === true ||
        status.Social === true ||
        status.children === true
      ) {
        return res.json({ status: true });
      }
      return res.status(400).json({ message: 'Dados informados vazio' });
    } catch (error) {
      console.log(error);
      return res.status(400).json({ status: 'error' });
    }
  }
}

export default new ViagemController();
