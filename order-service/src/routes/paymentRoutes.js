// File: src/routes/paymentRoutes.js
const express = require('express');
const router = express.Router();
const Payment = require('../models/Payment');
const Order = require('../models/Order');

/**
 * @swagger
 * /api/payments/process:
 *   post:
 *     summary: Process payment for an order
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               orderId:
 *                 type: string
 *               paymentMethod:
 *                 type: string
 *               paymentDetails:
 *                 type: object
 *     responses:
 *       200:
 *         description: Payment processed successfully
 *       404:
 *         description: Order not found
 */
router.post('/process', async (req, res) => {
  try {
    const { orderId, paymentMethod, paymentDetails } = req.body;
    
    if (!orderId || !paymentMethod) {
      return res.status(400).json({ error: true, message: 'Missing required fields' });
    }
    
    // Find the order
    const order = await Order.findById(orderId);
    
    if (!order) {
      return res.status(404).json({ error: true, message: 'Order not found' });
    }
    
    // Simulate payment processing
    const isPaymentSuccessful = Math.random() > 0.1; // 90% success rate for simulation
    
    // Create payment record
    const payment = new Payment({
      orderId: order._id,
      amount: order.totalAmount,
      method: paymentMethod,
      status: isPaymentSuccessful ? 'completed' : 'failed',
      paymentDetails: paymentDetails || {}
    });
    
    await payment.save();
    
    // Update order payment status
    order.paymentInfo.status = isPaymentSuccessful ? 'completed' : 'failed';
    order.paymentInfo.transactionId = payment.transactionId;
    
    // Update order status if payment is successful
    if (isPaymentSuccessful) {
      order.status = 'processing';
    }
    
    await order.save();
    
    res.status(200).json({
      success: isPaymentSuccessful,
      paymentId: payment._id,
      transactionId: payment.transactionId,
      status: payment.status,
      order: order
    });
  } catch (error) {
    console.error('Error processing payment:', error);
    res.status(500).json({ error: true, message: 'Failed to process payment' });
  }
});

/**
 * @swagger
 * /api/payments/{paymentId}:
 *   get:
 *     summary: Get payment by ID
 *     parameters:
 *       - in: path
 *         name: paymentId
 *         required: true
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Payment retrieved successfully
 *       404:
 *         description: Payment not found
 */
router.get('/:paymentId', async (req, res) => {
  try {
    const { paymentId } = req.params;
    
    const payment = await Payment.findById(paymentId);
    
    if (!payment) {
      return res.status(404).json({ error: true, message: 'Payment not found' });
    }
    
    res.status(200).json(payment);
  } catch (error) {
    console.error('Error fetching payment:', error);
    res.status(500).json({ error: true, message: 'Failed to fetch payment' });
  }
});

/**
 * @swagger
 * /api/payments/order/{orderId}:
 *   get:
 *     summary: Get payments by order ID
 *     parameters:
 *       - in: path
 *         name: orderId
 *         required: true
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Payments retrieved successfully
 */
router.get('/order/:orderId', async (req, res) => {
  try {
    const { orderId } = req.params;
    
    const payments = await Payment.find({ orderId }).sort({ createdAt: -1 });
    
    res.status(200).json(payments);
  } catch (error) {
    console.error('Error fetching payments by order:', error);
    res.status(500).json({ error: true, message: 'Failed to fetch payments' });
  }
});

module.exports = router;