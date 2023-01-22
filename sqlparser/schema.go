package sqlparser

// Define enums
type RelationshipType int
const (
    OneToOne RelationshipType = iota
    OneToMany
    ManyToMany
)

type AttributeType int
const (
    Varchar AttributeType = iota
    Number
    Other
)

// Define types
type Schema struct {
    Name        string
    Entities    []*Entity
}

type Entity struct {
    Name            string
    Cardinalities   []*Cardinality
    Attributes      []*Attribute
}

type Cardinality struct {
    MToN            [2]int
    Entity          *Entity
    Relationship    *Relationship
}

type Relationship struct {
    Type            *RelationshipType
    Cardinalities   *Cardinality
    Attributes      *Attribute
}

type Attribute struct {
    Name    string
    Type    AttributeType
}
