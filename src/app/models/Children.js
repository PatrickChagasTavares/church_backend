import Sequelize, { Model } from 'sequelize';

class Children extends Model {
  static init(sequelize) {
    super.init(
      {
        date: Sequelize.DATE,
        total: Sequelize.INTEGER,
        note: Sequelize.STRING,
      },
      { sequelize }
    );

    return this;
  }
}

export default Children;
