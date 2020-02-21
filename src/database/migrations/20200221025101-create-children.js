module.exports = {
  up: (queryInterface, Sequelize) => {
    return queryInterface.createTable('children', {
      id: {
        type: Sequelize.INTEGER,
        allowNull: false,
        autoIncrement: true,
        primaryKey: true,
      },
      date: { type: Sequelize.DATE, allowNull: false },
      total: { type: Sequelize.INTEGER, allowNull: false },
      note: { type: Sequelize.STRING },
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
    return queryInterface.dropTable('children');
  },
};
