require 'hanami-validations'

class User
  include Hanami::Validations

  validations do
    required(:name) { filled? }
  end
end
