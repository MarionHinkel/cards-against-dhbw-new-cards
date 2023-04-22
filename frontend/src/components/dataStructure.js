const CardColor = {
    BLACK: "black",
    WHITE: "white",
    GOLDEN: "golden"
}

class Player {
    constructor(id, name, isMod, isBoss, points) {
        this.id = id
        this.name = name
        this.isMod = isMod
        this.isBoss = isBoss
        this.points = points
    }
}

class GameCard {
    constructor(color, content, id) {
        this.color = color
        this.id = id
        this.content = content
        this.owner = ''
    }

    setOwner(owner) {
        this.owner = owner;
    }

}

export { CardColor, Player, GameCard }
