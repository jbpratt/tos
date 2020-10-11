import SwiftUI


struct Unwrap<Value, Content: View>: View {
    private let value: Value?
    private let contentProvider: (Value) -> Content

    init(_ value: Value?,
         @ViewBuilder content: @escaping (Value) -> Content)
    {
        self.value = value
        contentProvider = content
    }

    var body: some View {
        value.map(contentProvider)
    }
}
