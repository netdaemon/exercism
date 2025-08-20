public class SalaryCalculator {
    private static final double BASE_SALARY = 1000.00;
    private static final double MAX_SALARY = 2000.00;
    
    public double salaryMultiplier(int daysSkipped) {
        return daysSkipped < 5
            ? 1.0
            : 0.85;
    }

    public int bonusMultiplier(int productsSold) {
        return productsSold < 20
            ? 10
            : 13;
    }

    public double bonusForProductsSold(int productsSold) {
        int multiplier = bonusMultiplier(productsSold);

        return productsSold * multiplier;
    }

    public double finalSalary(int daysSkipped, int productsSold) {
        double multiplier = salaryMultiplier(daysSkipped);
        double bonus = bonusForProductsSold(productsSold);
        double salary = BASE_SALARY * multiplier + bonus;
        

        return salary > MAX_SALARY 
            ? MAX_SALARY 
            : salary; 
    } 
}
