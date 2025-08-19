class Badge {
    public String print(Integer id, String name, String department) {
        String badgeName = "";

        if (id != null) {
            badgeName += "[" + id + "] - ";
        }

        badgeName += name + " - ";

        if (department == null) {
            badgeName += "OWNER";
        } else {
            badgeName += department.toUpperCase();
        }

        return badgeName;
    }
}
