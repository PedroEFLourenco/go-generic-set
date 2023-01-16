
func TestSet_Contains(t *testing.T) {
 testSet := NewSetOfType[string]()
 stringA := "a"

 testSet.Add(stringA)

 assert.Equal(
  t,
  true,
  testSet.Contains(stringA),
  "TestSet should contain entry as expected.",
 )
}

func TestSet_Add(t *testing.T) {
 testSet := NewSetOfType[string]()
 stringA := "a"

 testSet.Add(stringA)

 assert.Equal(
  t,
  true,
  testSet.Contains(stringA),
  "TestSet should contain inserted entry.",
 )
}

func TestSet_ContainsAll(t *testing.T) {
 testSet := NewSetOfType[string]()
 stringA := "a"
 stringB := "b"
 stringC := "c"

 testEntries := []string{stringA, stringB, stringC}

 testSet.AddAll(testEntries)

 assert.Equal(
  t,
  true,
  testSet.ContainsAll(testEntries),
  "The Set should contain all of the provided entries",
 )
}

func TestSet_AddAll(t *testing.T) {
 testSet := NewSetOfType[string]()
 stringA := "a"
 stringB := "b"
 stringC := "c"

 testEntries := []string{stringA, stringB, stringC}

 testSet.AddAll(testEntries)

 assert.Equal(
  t,
  len(testEntries),
  testSet.GetSize(),
  "The Set should have the same size as the provided entries slice",
 )

 assert.Equal(
  t,
  true,
  testSet.ContainsAll(testEntries),
  "The Set should contain all of the provided entries",
 )
}

func TestSet_GetSize(t *testing.T) {
 testSet := NewSetOfType[string]()
 stringA := "a"
 stringB := "b"
 stringC := "c"

 testEntries := []string{stringA, stringB, stringC}

 for _, entry := range testEntries {
  testSet.Add(entry)
 }

 assert.Equal(
  t,
  len(testEntries),
  testSet.GetSize(),
  "TestSet size should match the number of provided entries",
 )
}

func TestSet_Remove(t *testing.T) {
 testSet := NewSetOfType[string]()
 stringA := "a"
 stringB := "b"
 stringC := "c"

 testEntries := []string{stringA, stringB, stringC}

 testSet.AddAll(testEntries)
 testSet.Remove(stringA)

 assert.Equal(
  t,
  false,
  testSet.Contains(stringA),
  "TestSet should not contain removed entry",
 )

 assert.Equal(
  t,
  len(testEntries)-1,
  testSet.GetSize(),
  "TestSet size should match the number of provided entries minus the removed one",
 )
}

func TestSet_RemoveAll(t *testing.T) {
 testSet := NewSetOfType[string]()
 stringA := "a"
 stringB := "b"
 stringC := "c"

 testEntries := []string{stringA, stringB, stringC}

 testSet.AddAll(testEntries)
 testSet.RemoveAll(testEntries)

 assert.Equal(
  t,
  0,
  testSet.GetSize(),
  "TestSet size should be zero since all entries were removed",
 )
}

func TestSet_ToSlice(t *testing.T) {
 testSet := NewSetOfType[string]()
 stringA := "a"
 stringB := "b"
 stringC := "c"

 testEntries := []string{stringA, stringB, stringC}

 testSet.AddAll(testEntries)

 resultingSlice := testSet.ToSlice()

 assert.Equal(
  t,
  len(testEntries),
  len(resultingSlice),
  "Slice size should match the number of provided entries in the set",
 )

 assert.True(
  t,
  testSet.ContainsAll(testEntries),
  "Slice should contain all entries from the set",
 )
}

func TestSet_ExtractMapPrimitive(t *testing.T) {
 testSet := NewSetOfType[string]()
 mapPrimitive := testSet.ExtractMapPrimitive()

 assert.Equal(
  t,
  "map[string]bool",
  reflect.TypeOf(mapPrimitive).String(),
  "result should be of the map and it isn't",
 )
}
