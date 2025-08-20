class SqueakyClean {
    static String clean(String identifier) {
        char[] characters = identifier.toCharArray();
        StringBuilder sb = new StringBuilder();

        boolean upCaseNextChar = false;
        
        for (int i = 0; i < characters.length; i++) {
            char character = characters[i];
            
            if (character == ' ') {
                sb.append('_');
            } else if (character == '-') {
                upCaseNextChar = true;
            } else if(character == '4') {
                sb.append('a');
            } else if(character == '3') {
                sb.append('e');
            } else if(character == '0') {
                sb.append('o');
            } else if(character == '1') {
                sb.append('l');
            } else if(character == '7') {
                sb.append('t');
            } else if (Character.isLetter(character)){
                if (upCaseNextChar) {
                    sb.append(Character.toUpperCase(character));
                    upCaseNextChar = false;
                } else {
                    sb.append(character);                
                }
            }
        }

        return sb.toString();
    }
}
