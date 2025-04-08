// auth-service/src/config.js
module.exports = {
  mongoURI: process.env.MONGO_URI || 'mongodb://mongo:27017/auth',
  jwtSecret: process.env.JWT_SECRET || 'your_jwt_secret_key_here',
  jwtExpiresIn: process.env.JWT_EXPIRES_IN || '1d',
  saltRounds: 10,
  roles: {
    CUSTOMER: 'customer',
    ADMIN: 'admin'
  }
};
