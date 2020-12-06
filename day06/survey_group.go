package day06

// A SurveyGroup represents a collection of answers to a yes/no questionnaire
// by a grouop of people.
type SurveyGroup struct {
	// memberCount is the number of people in this survey group.
	memberCount int
	// tally is a map from a question's identifying letter to the count of
	// members who responded 'yes' to that question.
	tally map[rune]int
}

func makeSurveyGroup() SurveyGroup {
	return SurveyGroup{
		tally: make(map[rune]int),
	}
}

// ParseLinesToSurveyGroups takes a slice of strings and makes a slice of
// SurveyGroups. Each line represents a person and the letters of the questions
// they answered 'yes' to. Groups are separated by blank lines.
func ParseLinesToSurveyGroups(lines []string) []SurveyGroup {
	var groups []SurveyGroup

	sg := makeSurveyGroup()
	for _, line := range lines {
		// Groups are separated by blank lines. Append current group and
		// create a new empty group for the next iteration.
		if line == "" {
			groups = append(groups, sg)
			sg = makeSurveyGroup()

			continue
		}

		sg.memberCount++

		// Add this person's responses to the tally.
		for _, r := range line {
			if _, ok := sg.tally[r]; !ok {
				sg.tally[r] = 0
			}
			sg.tally[r]++
		}
	}

	// If the survey group has data inside it at the end of the loop, we shall
	// append it to the list as well. This catches the last group.
	if sg.memberCount > 0 {
		groups = append(groups, sg)
	}

	return groups
}

// FindCommonQuestionCount finds the number of questions that ALL members in the
// survey group responded 'yes' to.
func (sg *SurveyGroup) FindCommonQuestionCount() int {
	var count int
	for _, n := range sg.tally {
		if n == sg.memberCount {
			count++
		}
	}
	return count
}

// FindAnsweredQuestionCount finds the number of questions that ANY member of
// the survey group responded 'yes' to.
func (sg *SurveyGroup) FindAnsweredQuestionCount() int { return len(sg.tally) }
