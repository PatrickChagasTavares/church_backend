import Sequelize, { Model } from 'sequelize';

class Social extends Model {
  static init(sequelize) {
    super.init(
      {
        date: Sequelize.DATE,
        nameTribe: Sequelize.STRING,
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
        note: Sequelize.STRING,
      },
      { sequelize }
    );

    return this;
  }
}

export default Social;
