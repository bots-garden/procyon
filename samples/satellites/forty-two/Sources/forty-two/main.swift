import Suborbital

class FortyTwo: Suborbital.Runnable {
    func run(input: String) -> String {
        return """
        {"value":42}
        """
    }
}

Suborbital.Set(runnable: FortyTwo())
