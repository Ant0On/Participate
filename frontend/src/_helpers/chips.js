const chipsMapper = (discount = 0) => {
    return {
    "sports_event": {
        color: "deep-purple",
        icon: "mdi-stadium",
        text: "Sports event"
    },
    "festival": {
        color: "pink-darken-2",
        icon: "mdi-account-group",
        text: "Festival"
    },
    "concert": {
        color: "amber",
        icon: "mdi-music",
        text: "Concert"
    },
    "conference": {
        color: "indigo",
        icon: "mdi-library",
        text: "Conference"
    },
    "beginner": {
        color: "green",
        icon: "mdi-yoga",
        text: "Beginner"
    },
    "intermediate": {
        color: "orange",
        icon: "mdi-bullseye",
        text: "Intermediate"
    },
    "advanced": {
        color: "red",
        icon: "mdi-hiking",
        text: "Advanced"
    },
    "indoor": {
        color: "grey",
        icon: "mdi-home",
        text: "Indoor"
    },
    "outdoor": {
        color: "blue-lighten-1",
        icon: "mdi-cloud",
        text: "Outdoor"
    },
    "animal_friendly": {
        color: "blue",
        icon: "mdi-dog-side",
        text: "Animal friendly"
    },
    "guesthouse": {
        color: "light-blue-darken-4",
        icon: "mdi-home",
        text: "Guesthouse"
    },
    "villa": {
        color: "deep-purple-accent-4",
        icon: "mdi-warehouse",
        text: "Villa"
    },
    "apartment": {
        color: "green-accent-4",
        icon: "mdi-city-variant",
        text: "Apartment"
    },
    "hostel": {
        color: "blue-darken-4",
        icon: "mdi-office-building",
        text: "Hostel"
    },
    "hotel": {
        color: "deep-orange",
        icon: "mdi-domain",
        text: "Hotel"
    },
    "recommended": {
        color: "amber",
        icon: "mdi-star-david",
        text: "Recommended"
    },
    "discount": {
        color: "red",
        icon: "mdi-tag",
        text: `-${discount || 0}%`
    }
}
}
export default chipsMapper