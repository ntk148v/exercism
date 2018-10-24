/* twofer or 2fer is short for two for one. One for you and one for me.

"One for X, one for me."

When X is a name or "you".

If the given name is "Kien", the result should be "One for Kien, one for me."
If no name is given, the result should be "One for you, one for me."
*/
package twofer

// ShareWith apportions things between two people
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
