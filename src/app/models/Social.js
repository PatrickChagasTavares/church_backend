import Sequelize, { Model } from 'sequelize';

class Social extends Model {
  static init(sequelize) {
    super.init(
      {
        date: Sequelize.DATE,
        namePerson: Sequelize.STRING,
        address: Sequelize.STRING,
        type: Sequelize.STRING,
        age: Sequelize.STRING,
        bible: Sequelize.STRING,
        evangelical: Sequelize.STRING,
        contact: Sequelize.STRING,
        frequentsChurch: Sequelize.STRING,
        cultHome: Sequelize.STRING,
        studyBible: Sequelize.STRING,
        studyConfimation: Sequelize.STRING,
        prayerRequest: Sequelize.STRING,
        reconciled: Sequelize.STRING,
        visit: Sequelize.STRING,
        acceptChrist: Sequelize.STRING,
        medical: Sequelize.STRING,
        optician: Sequelize.STRING,
        pressure: Sequelize.STRING,
        glucose: Sequelize.STRING,
        aesthetics: Sequelize.STRING,
        cuttingHair: Sequelize.STRING,
        hairstyle: Sequelize.STRING,
        Nail: Sequelize.STRING,
        Eyebrow: Sequelize.STRING,
        note: Sequelize.STRING,
      },
      { sequelize }
    );

    return this;
  }
}

export default Social;
