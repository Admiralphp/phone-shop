// auth-service/src/swagger.js
const swaggerJsDoc = require('swagger-jsdoc');

// Swagger definition
const swaggerDefinition = {
  openapi: '3.0.0',
  info: {
    title: 'Authentication API',
    version: '1.0.0',
    description: 'Authentication microservice for phone accessories web application',
    contact: {
      name: 'API Support',
      email: 'support@phoneaccessories.com'
    }
  },
  servers: [
    {
      url: 'http://localhost:3000',
      description: 'Development server'
    }
  ],
  components: {
    securitySchemes: {
      bearerAuth: {
        type: 'http',
        scheme: 'bearer',
        bearerFormat: 'JWT'
      }
    },
    schemas: {
      User: {
        type: 'object',
        properties: {
          id: {
            type: 'string',
            description: 'User ID'
          },
          email: {
            type: 'string',
            description: 'User email'
          },
          firstName: {
            type: 'string',
            description: 'User first name'
          },
          lastName: {
            type: 'string',
            description: 'User last name'
          },
          role: {
            type: 'string',
            enum: ['customer', 'admin'],
            description: 'User role'
          }
        }
      },
      RegisterRequest: {
        type: 'object',
        required: ['email', 'password', 'firstName', 'lastName'],
        properties: {
          email: {
            type: 'string',
            format: 'email',
            description: 'User email'
          },
          password: {
            type: 'string',
            format: 'password',
            minLength: 6,
            description: 'User password'
          },
          firstName: {
            type: 'string',
            description: 'User first name'
          },
          lastName: {
            type: 'string',
            description: 'User last name'
          },
          role: {
            type: 'string',
            enum: ['customer', 'admin'],
            description: 'User role (defaults to customer if not provided)'
          }
        }
      },
      LoginRequest: {
        type: 'object',
        required: ['email', 'password'],
        properties: {
          email: {
            type: 'string',
            format: 'email',
            description: 'User email'
          },
          password: {
            type: 'string',
            format: 'password',
            description: 'User password'
          }
        }
      },
      UpdateProfileRequest: {
        type: 'object',
        properties: {
          firstName: {
            type: 'string',
            description: 'User first name'
          },
          lastName: {
            type: 'string',
            description: 'User last name'
          }
        }
      },
      AuthResponse: {
        type: 'object',
        properties: {
          success: {
            type: 'boolean',
            description: 'Operation result'
          },
          message: {
            type: 'string',
            description: 'Response message'
          },
          data: {
            type: 'object',
            properties: {
              token: {
                type: 'string',
                description: 'JWT token'
              },
              user: {
                $ref: '#/components/schemas/User'
              }
            }
          }
        }
      },
      ErrorResponse: {
        type: 'object',
        properties: {
          success: {
            type: 'boolean',
            description: 'Operation result',
            example: false
          },
          message: {
            type: 'string',
            description: 'Error message'
          },
          error: {
            type: 'string',
            description: 'Detailed error information'
          }
        }
      }
    }
  }
};

// Options for the swagger docs
const options = {
  swaggerDefinition,
  // Paths to files containing OpenAPI definitions
  apis: ['./src/routes.js']
};

// Initialize swagger-jsdoc
const swaggerSpec = swaggerJsDoc(options);

module.exports = swaggerSpec;
