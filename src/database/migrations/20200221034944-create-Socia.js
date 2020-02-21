module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable('social', {
      id: {
        type: Sequelize.INTEGER,
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
      },
      date: {
        type: Sequelize.DATE,
        allowNull: false,
      },
      nameTribe: {
        type: Sequelize.STRING,
        allowNull: false,
      },
      namePerson: {
        type: Sequelize.STRING,
        allowNull: false,
      },
      address: {
        type: Sequelize.STRING,
        allowNull: false,
      },
      type: {
        type: Sequelize.STRING(20),
        allowNull: false,
      },
      age: {
        type: Sequelize.STRING(20),
        allowNull: false,
      },
      bible: {
        type: Sequelize.STRING(20),
      },
      evangelical: {
        type: Sequelize.STRING(20),
      },
      contact: {
        type: Sequelize.STRING(20),
      },
      frequentsChurch: {
        type: Sequelize.STRING(20),
      },
      cultHome: {
        type: Sequelize.STRING(20),
      },
      studyBible: {
        type: Sequelize.STRING(20),
      },
      studyConfimation: {
        type: Sequelize.STRING(20),
      },
      prayerRequest: {
        type: Sequelize.STRING(20),
      },
      reconciled: {
        type: Sequelize.STRING(20),
      },
      visit: {
        type: Sequelize.STRING(20),
      },
      acceptChrist: {
        type: Sequelize.STRING(20),
      },
      medical: {
        type: Sequelize.STRING(20),
      },
      optician: {
        type: Sequelize.STRING(20),
      },
      pressure: {
        type: Sequelize.STRING(20),
      },
      glucose: {
        type: Sequelize.STRING(20),
      },
      aesthetics: {
        type: Sequelize.STRING(20),
      },
      cuttingHair: {
        type: Sequelize.STRING(20),
      },
      hairstyle: {
        type: Sequelize.STRING(20),
      },
      Nail: {
        type: Sequelize.STRING(20),
      },
      Eyebrow: {
        type: Sequelize.STRING(20),
      },
      note: {
        type: Sequelize.STRING,
      },
      created_at: {
        type: Sequelize.DATE,
        allowNull: false,
      },
      updated_at: {
        type: Sequelize.DATE,
        allowNull: false,
      },
    });
  },

  down: queryInterface => {
    return queryInterface.dropTable('social');
  },
};
