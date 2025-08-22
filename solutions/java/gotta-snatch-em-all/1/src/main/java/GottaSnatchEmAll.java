import java.util.List;
import java.util.Set;
import java.util.HashSet;

class GottaSnatchEmAll {

    static Set<String> newCollection(List<String> cards) {
        Set<String> collection = new HashSet<String>();
        collection.addAll(cards);
        
        return collection;
    }

    static boolean addCard(String card, Set<String> collection) {
        return collection.add(card);
    }

    static boolean canTrade(Set<String> myCollection, Set<String> theirCollection) {
        boolean theyHaveUniq = myCollection.stream()
            .anyMatch(c -> !theirCollection.contains(c));
        boolean iHaveUniq = theirCollection.stream()
            .anyMatch(c -> !myCollection.contains(c));

        return theyHaveUniq && iHaveUniq;
    }

    static Set<String> commonCards(List<Set<String>> collections) {
        Set<String> common = new HashSet<>(collections.get(0));

        for (Set<String> collection : collections) {
            common.retainAll(collection);
        }

        return common;
    }

    static Set<String> allCards(List<Set<String>> collections) {
        Set<String> allCards =  new HashSet<>();

        for (Set<String> collection : collections) {
            allCards.addAll(collection);
        }

        return allCards;
    }
}
