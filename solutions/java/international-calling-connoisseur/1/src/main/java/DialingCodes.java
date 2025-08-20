import java.util.HashMap;
import java.util.Map;


public class DialingCodes {
    private Map<Integer, String> countryCodes = new HashMap<>();

    public Map<Integer, String> getCodes() {
        return countryCodes;
    }

    public void setDialingCode(Integer code, String country) {
        countryCodes.put(code, country);
    }

    public String getCountry(Integer code) {
        return countryCodes.get(code);
    }

    public void addNewDialingCode(Integer code, String country) {
        if (countryCodes.containsKey(code) || countryCodes.containsValue(country)) {
            return;
        }
        
        countryCodes.put(code, country);
    }

    public Integer findDialingCode(String country) {
        for (Integer key : countryCodes.keySet()) {
            if (countryCodes.get(key) == country) {
                return key;
            }
        }

        return null;
    }

    public void updateCountryDialingCode(Integer code, String country) {
        Integer dialingCode = findDialingCode(country);

        if (dialingCode != null) {
            countryCodes.remove(dialingCode);
        }
        
        countryCodes.put(code, country);
    }
}