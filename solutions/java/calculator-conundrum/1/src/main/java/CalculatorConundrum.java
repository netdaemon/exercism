class CalculatorConundrum {
    public String calculate(int operand1, int operand2, String operation) {
        switch(operation) {
            case "+":
                return add(operand1, operand2);
            case "-":
                return subtract(operand1, operand2);
            case "*":
                return multiply(operand1, operand2);
            case "/":
                return divide(operand1, operand2);
            case null:
                throw new IllegalArgumentException("Operation cannot be null");
            case "":
                throw new IllegalArgumentException("Operation cannot be empty");
            default:
                throw new IllegalOperationException("Operation '" + operation + "' does not exist");
        }
    }

    String add(int operand1, int operand2) {
        int result = operand1 + operand2;
        return operand1 + " + "  + operand2 + " = " + result;
    }

    String subtract(int operand1, int operand2) {
        int result = operand1 - operand2;
        return operand1 + " - " + operand2 + " = " + result;
    }

    String multiply(int operand1, int operand2) {
        int result = operand1 * operand2;
        return operand1 +  " * " + operand2 + " = " + result;
    }

    String divide(int operand1, int operand2) {
        try {
            int result = operand1 / operand2;
            return operand1 +  " / " + operand2 + " = " + result;
        } catch(ArithmeticException e) {
            throw new IllegalOperationException("Division by zero is not allowed", e);
        }
        
        
    }
}
